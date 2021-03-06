import { equals } from 'ramda';
import { Observable } from 'rxjs';
import { map, filter, tap, share } from 'rxjs/operators';

import { Contact } from '@models';
import { SocketLayer } from '@layers/socket-layer/socket-layer';
import { Message } from '@protobuf/Message_pb';
import logger from '@utils/logging';

export class MessageLayer {
  private messages$: Observable<Message>;
  private receivedMessages$: Observable<Buffer>;

  constructor(private worker: SocketLayer) {
    this.receivedMessages$ = this.worker.getReceivedMessagesStream();
    this.messages$ = this.handleReceivedMessages();
  }

  close() {
    logger.info('Message layer: closing.');
    this.worker.close();
  }

  on(type: Message.MessageType): Observable<Message> {
    return this.messages$.pipe(filter((msg) => msg.getType() === type));
  }

  send(msg: Message): Promise<void> {
    const sender = msg.getSender();
    const receiver = msg.getReceiver();

    if (!sender || !receiver) {
      throw new Error('Message layer: Invalid message sender/receiver.');
    }

    const { address: senderAddress } = Contact.from(sender);
    const { address } = Contact.from(receiver);

    const msgToSelf = equals(senderAddress, address);
    if (msgToSelf) {
      throw new Error('Message layer: Cannot send message to self.');
    }

    logger.info('Message layer: encoding message.');
    return this.worker.send({ data: encodeMessage(msg), address });
  }

  private handleReceivedMessages() {
    return this.receivedMessages$.pipe(
      tap(() => logger.info('Message layer: decoding message.')),
      map(decodeMessage),
      share()
    );
  }
}

export function decodeMessage(buffer: Buffer): Message {
  return Message.deserializeBinary(Uint8Array.from(buffer));
}

export function encodeMessage(msg: Message): Buffer {
  return Buffer.from(msg.serializeBinary());
}

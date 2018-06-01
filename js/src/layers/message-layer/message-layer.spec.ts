import { Subject } from 'rxjs';

import { Communication, Contact, Address } from '@models';
import { Message } from '@protobuf/Message_pb';

import { MessageLayer, encodeMessage } from './message-layer';
import { prepareFindNodeMessage } from '@protobuf/utils';

describe('Layer: MessageLayer', function() {
  let inputMessages$: Subject<Buffer>;
  let outputMessages$: Subject<Communication<Buffer>>;
  let layer: MessageLayer;

  let senderAddress: Address;
  let receiverAddress: Address;
  let msg: Message;

  beforeEach(function() {
    inputMessages$ = new Subject();
    outputMessages$ = fakeSubject<Communication<Buffer>>('outputMessages');
    layer = new MessageLayer(inputMessages$.asObservable(), outputMessages$);

    senderAddress = { host: 'foo', port: 123 };
    receiverAddress = { host: 'bar', port: 234 };
    msg = prepareFindNodeMessage({
      node: 'foo',
      sender: new Contact({ address: senderAddress }),
      receiver: new Contact({
        address: receiverAddress
      })
    });
  });

  describe('Method: "on"', function() {
    it('should exist', function() {
      const result = typeof layer.on === 'function';
      expect(result).toBe(true);
    });

    describe('When: message of type "FIND_NODE" is in input stream', function() {
      let firstCallback: jasmine.Spy;
      let secondCallback: jasmine.Spy;
      let otherTypeCallback: jasmine.Spy;

      beforeEach(function() {
        firstCallback = jasmine.createSpy('firstCallback');
        secondCallback = jasmine.createSpy('secondCallback');
        otherTypeCallback = jasmine.createSpy('otherTypeCallback');
      });

      it('should invoke callbacks listening on that type', function() {
        layer.on(Message.MessageType.FIND_NODE).subscribe(firstCallback);
        layer.on(Message.MessageType.FIND_NODE).subscribe(secondCallback);

        inputMessages$.next(encodeMessage(msg));

        expect(firstCallback).toHaveBeenCalled();
        expect(secondCallback).toHaveBeenCalled();
      });

      it('should not invoke callback listening on other type', function() {
        layer.on(Message.MessageType.PING).subscribe(otherTypeCallback);

        inputMessages$.next(encodeMessage(msg));

        expect(otherTypeCallback).not.toHaveBeenCalled();
      });
    });
  });

  describe('Method: "send"', function() {
    it('should exist', function() {
      const result = typeof layer.on === 'function';
      expect(result).toBe(true);
    });

    describe('When: it is called with message', function() {
      beforeEach(function() {
        layer.send(msg);
      });

      it('output stream should contain valid data', function() {
        const expected = {
          data: encodeMessage(msg),
          address: receiverAddress
        };
        expect(outputMessages$.next).toHaveBeenCalledWith(expected);
      });
    });
  });
});

function fakeSubject<T>(name: string): Subject<T> {
  return {
    next: jasmine.createSpy(`${name}NextSpy`) as Function,
    error: jasmine.createSpy(`${name}ErrorSpy`) as Function,
    complete: jasmine.createSpy(`${name}CompleteSpy`) as Function
  } as Subject<T>;
}

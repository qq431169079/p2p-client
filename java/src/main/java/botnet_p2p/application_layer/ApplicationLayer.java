package botnet_p2p.application_layer;

import botnet_p2p.business_logic_layer.BusinessLogicLayer;
import botnet_p2p.model.Peer;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class ApplicationLayer {
    private static final Logger logger = LogManager.getLogger(ApplicationLayer.class);

    private BusinessLogicLayer businessLogicLayer;

    public ApplicationLayer(BusinessLogicLayer businessLogicLayer) {
        this.businessLogicLayer = businessLogicLayer;
    }

    public void launchClient(Peer bootstrapNode) {
        businessLogicLayer.joinNetwork(bootstrapNode);

        logger.info("bootstrap finished");
        readCommands();
    }

    public void startWithoutBootstrapping() {
        businessLogicLayer.createNetwork();

        readCommands();
    }

    public void shutdown() {
        logger.info("closing");
        businessLogicLayer.shutdown();
    }

    private void printRoutingTable() {
        logger.info(
                businessLogicLayer.getRoutingTable()
        );
    }

    private void readCommands() {
        while (true) {
            String command = System.console().readLine();
            if("p".equals(command)) {
                this.printRoutingTable();
            }
        }
    }
}

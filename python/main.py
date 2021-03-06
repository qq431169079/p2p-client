import sys
sys.path.append('../')
from python.Socket.SocketLayer import SocketLayer
from python.P2P.P2PLayer import P2PLayer, Peer
from python.Message.MessageLayer import MessageLayer
from python.Business.BusinessLogicLayer import BusinessLogicLayer
from python.Application.ApplicationLayer import Application
import asyncio
import logging.handlers

logging.basicConfig(
    level=logging.DEBUG,
    format='%(name)s: %(message)s',
)
handler = logging.handlers.RotatingFileHandler(
    "./logs/log.txt"
)

log = logging.getLogger(__name__)
log.addHandler(handler)
log.propagate = False
formatter = logging.Formatter('%(name)s: %(message)s')
handler.formatter = formatter

def _run(cor):
    return asyncio.get_event_loop().run_until_complete(cor)


def set_up_layers(ip, port):
    socket_layer = SocketLayer()
    message_layer = MessageLayer(socket_layer)
    p2p_layer = P2PLayer(message_layer, ip, port)
    business_logic_layer = BusinessLogicLayer(p2p_layer)
    application_layer = Application(business_logic_layer)
    return (socket_layer, message_layer, p2p_layer, business_logic_layer, application_layer)


def set_up_layers_communications(socket_layer, message_layer, p2p_layer, business_logic_layer):
    mess_to_sock_queue = asyncio.Queue()
    sock_to_mess_queue = asyncio.Queue()
    mess_to_p2p_queue = asyncio.Queue()
    p2p_to_mess_queue = asyncio.Queue()
    bll_to_p2p_queue = asyncio.Queue()
    p2p_to_bll_queue = asyncio.Queue()
    _run(socket_layer.add_layer_communication(higher=(mess_to_sock_queue, sock_to_mess_queue)))
    _run(message_layer.add_layer_communication(higher=(p2p_to_mess_queue, mess_to_p2p_queue),
                                               lower=(sock_to_mess_queue, mess_to_sock_queue)))
    _run(p2p_layer.add_layer_communication(higher=(bll_to_p2p_queue, p2p_to_bll_queue),
                                           lower=(mess_to_p2p_queue, p2p_to_mess_queue)))
    _run(business_logic_layer.add_layer_communication(lower=(p2p_to_bll_queue, bll_to_p2p_queue)))


if __name__ == '__main__':
    ip = sys.argv[1]
    port = int(sys.argv[2])

    layers = set_up_layers(ip, port)
    set_up_layers_communications(*layers[0:-1])

    app = layers[-1]
    app.run()
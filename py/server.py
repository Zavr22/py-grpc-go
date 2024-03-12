from concurrent.futures import ThreadPoolExecutor

import grpc

from hello_pb2_grpc import HelloServicer, add_HelloServicer_to_server


class HelloServer(HelloServicer):
    def Hello(self, request, context):
        return str(request.data + "sosi")

    if __name__ == '__main__':
        server = grpc.server(ThreadPoolExecutor())
        add_HelloServicer_to_server(HelloServicer, server)
        port = 5436
        server.add_insecure_port(f'[::]:{port}')
        server.start()
        server.wait_for_termination()
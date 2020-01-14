# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import data_pb2 as data__pb2


class DataStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.FetchData = channel.unary_unary(
        '/proto.Data/FetchData',
        request_serializer=data__pb2.FetchDataRequest.SerializeToString,
        response_deserializer=data__pb2.FetchDataReponse.FromString,
        )
    self.LoadData = channel.unary_unary(
        '/proto.Data/LoadData',
        request_serializer=data__pb2.LoadDataRequest.SerializeToString,
        response_deserializer=data__pb2.LoadDataReponse.FromString,
        )
    self.TimerTask = channel.unary_unary(
        '/proto.Data/TimerTask',
        request_serializer=data__pb2.TimerTaskRequest.SerializeToString,
        response_deserializer=data__pb2.TimerTaskResponse.FromString,
        )
    self.GetSymbolOhlc = channel.unary_unary(
        '/proto.Data/GetSymbolOhlc',
        request_serializer=data__pb2.GetSymbolOhlcRequest.SerializeToString,
        response_deserializer=data__pb2.GetSymbolOhlcResponse.FromString,
        )
    self.PredictSymbol = channel.stream_stream(
        '/proto.Data/PredictSymbol',
        request_serializer=data__pb2.PredictSymbolRequest.SerializeToString,
        response_deserializer=data__pb2.PredictSymbolResponse.FromString,
        )


class DataServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def FetchData(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def LoadData(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def TimerTask(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetSymbolOhlc(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PredictSymbol(self, request_iterator, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_DataServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'FetchData': grpc.unary_unary_rpc_method_handler(
          servicer.FetchData,
          request_deserializer=data__pb2.FetchDataRequest.FromString,
          response_serializer=data__pb2.FetchDataReponse.SerializeToString,
      ),
      'LoadData': grpc.unary_unary_rpc_method_handler(
          servicer.LoadData,
          request_deserializer=data__pb2.LoadDataRequest.FromString,
          response_serializer=data__pb2.LoadDataReponse.SerializeToString,
      ),
      'TimerTask': grpc.unary_unary_rpc_method_handler(
          servicer.TimerTask,
          request_deserializer=data__pb2.TimerTaskRequest.FromString,
          response_serializer=data__pb2.TimerTaskResponse.SerializeToString,
      ),
      'GetSymbolOhlc': grpc.unary_unary_rpc_method_handler(
          servicer.GetSymbolOhlc,
          request_deserializer=data__pb2.GetSymbolOhlcRequest.FromString,
          response_serializer=data__pb2.GetSymbolOhlcResponse.SerializeToString,
      ),
      'PredictSymbol': grpc.stream_stream_rpc_method_handler(
          servicer.PredictSymbol,
          request_deserializer=data__pb2.PredictSymbolRequest.FromString,
          response_serializer=data__pb2.PredictSymbolResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'proto.Data', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))

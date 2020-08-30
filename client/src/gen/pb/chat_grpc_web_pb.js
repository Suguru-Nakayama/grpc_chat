/**
 * @fileoverview gRPC-Web generated client stub for chat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.chat = require('./chat_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.ChatClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.ChatPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chat.CreateChatRoomRequest,
 *   !proto.chat.CreateChatRoomResponse>}
 */
const methodDescriptor_Chat_CreateChatRoom = new grpc.web.MethodDescriptor(
  '/chat.Chat/CreateChatRoom',
  grpc.web.MethodType.UNARY,
  proto.chat.CreateChatRoomRequest,
  proto.chat.CreateChatRoomResponse,
  /**
   * @param {!proto.chat.CreateChatRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.CreateChatRoomResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chat.CreateChatRoomRequest,
 *   !proto.chat.CreateChatRoomResponse>}
 */
const methodInfo_Chat_CreateChatRoom = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chat.CreateChatRoomResponse,
  /**
   * @param {!proto.chat.CreateChatRoomRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.CreateChatRoomResponse.deserializeBinary
);


/**
 * @param {!proto.chat.CreateChatRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.chat.CreateChatRoomResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chat.CreateChatRoomResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.createChatRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chat.Chat/CreateChatRoom',
      request,
      metadata || {},
      methodDescriptor_Chat_CreateChatRoom,
      callback);
};


/**
 * @param {!proto.chat.CreateChatRoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chat.CreateChatRoomResponse>}
 *     A native promise that resolves to the response
 */
proto.chat.ChatPromiseClient.prototype.createChatRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/chat.Chat/CreateChatRoom',
      request,
      metadata || {},
      methodDescriptor_Chat_CreateChatRoom);
};


module.exports = proto.chat;


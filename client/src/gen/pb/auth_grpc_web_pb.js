/**
 * @fileoverview gRPC-Web generated client stub for auth
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.auth = require('./auth_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.auth.AuthClient =
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
proto.auth.AuthPromiseClient =
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
 *   !proto.auth.SignUpRequest,
 *   !proto.auth.SignUpResponse>}
 */
const methodDescriptor_Auth_SignUp = new grpc.web.MethodDescriptor(
  '/auth.Auth/SignUp',
  grpc.web.MethodType.UNARY,
  proto.auth.SignUpRequest,
  proto.auth.SignUpResponse,
  /**
   * @param {!proto.auth.SignUpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.SignUpResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.auth.SignUpRequest,
 *   !proto.auth.SignUpResponse>}
 */
const methodInfo_Auth_SignUp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.auth.SignUpResponse,
  /**
   * @param {!proto.auth.SignUpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.SignUpResponse.deserializeBinary
);


/**
 * @param {!proto.auth.SignUpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.auth.SignUpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.SignUpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthClient.prototype.signUp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.Auth/SignUp',
      request,
      metadata || {},
      methodDescriptor_Auth_SignUp,
      callback);
};


/**
 * @param {!proto.auth.SignUpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.SignUpResponse>}
 *     A native promise that resolves to the response
 */
proto.auth.AuthPromiseClient.prototype.signUp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.Auth/SignUp',
      request,
      metadata || {},
      methodDescriptor_Auth_SignUp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.LogInRequest,
 *   !proto.auth.LogInResponse>}
 */
const methodDescriptor_Auth_LogIn = new grpc.web.MethodDescriptor(
  '/auth.Auth/LogIn',
  grpc.web.MethodType.UNARY,
  proto.auth.LogInRequest,
  proto.auth.LogInResponse,
  /**
   * @param {!proto.auth.LogInRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.LogInResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.auth.LogInRequest,
 *   !proto.auth.LogInResponse>}
 */
const methodInfo_Auth_LogIn = new grpc.web.AbstractClientBase.MethodInfo(
  proto.auth.LogInResponse,
  /**
   * @param {!proto.auth.LogInRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.LogInResponse.deserializeBinary
);


/**
 * @param {!proto.auth.LogInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.auth.LogInResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.LogInResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthClient.prototype.logIn =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.Auth/LogIn',
      request,
      metadata || {},
      methodDescriptor_Auth_LogIn,
      callback);
};


/**
 * @param {!proto.auth.LogInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.LogInResponse>}
 *     A native promise that resolves to the response
 */
proto.auth.AuthPromiseClient.prototype.logIn =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.Auth/LogIn',
      request,
      metadata || {},
      methodDescriptor_Auth_LogIn);
};


module.exports = proto.auth;


package com.zgg2001.grpc.service;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.33.0)",
    comments = "Source: service.proto")
public final class ProjectServiceGrpc {

  private ProjectServiceGrpc() {}

  public static final String SERVICE_NAME = "com.zgg2001.grpc.service.ProjectService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.LPCheckRequest,
      com.zgg2001.grpc.service.Service.LPCheckResponse> getLicencePlateCheckMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LicencePlateCheck",
      requestType = com.zgg2001.grpc.service.Service.LPCheckRequest.class,
      responseType = com.zgg2001.grpc.service.Service.LPCheckResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.LPCheckRequest,
      com.zgg2001.grpc.service.Service.LPCheckResponse> getLicencePlateCheckMethod() {
    io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.LPCheckRequest, com.zgg2001.grpc.service.Service.LPCheckResponse> getLicencePlateCheckMethod;
    if ((getLicencePlateCheckMethod = ProjectServiceGrpc.getLicencePlateCheckMethod) == null) {
      synchronized (ProjectServiceGrpc.class) {
        if ((getLicencePlateCheckMethod = ProjectServiceGrpc.getLicencePlateCheckMethod) == null) {
          ProjectServiceGrpc.getLicencePlateCheckMethod = getLicencePlateCheckMethod =
              io.grpc.MethodDescriptor.<com.zgg2001.grpc.service.Service.LPCheckRequest, com.zgg2001.grpc.service.Service.LPCheckResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "LicencePlateCheck"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.LPCheckRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.LPCheckResponse.getDefaultInstance()))
              .build();
        }
      }
    }
    return getLicencePlateCheckMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UploadInfoRequest,
      com.zgg2001.grpc.service.Service.UploadInfoResponse> getUploadParkingInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UploadParkingInfo",
      requestType = com.zgg2001.grpc.service.Service.UploadInfoRequest.class,
      responseType = com.zgg2001.grpc.service.Service.UploadInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UploadInfoRequest,
      com.zgg2001.grpc.service.Service.UploadInfoResponse> getUploadParkingInfoMethod() {
    io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UploadInfoRequest, com.zgg2001.grpc.service.Service.UploadInfoResponse> getUploadParkingInfoMethod;
    if ((getUploadParkingInfoMethod = ProjectServiceGrpc.getUploadParkingInfoMethod) == null) {
      synchronized (ProjectServiceGrpc.class) {
        if ((getUploadParkingInfoMethod = ProjectServiceGrpc.getUploadParkingInfoMethod) == null) {
          ProjectServiceGrpc.getUploadParkingInfoMethod = getUploadParkingInfoMethod =
              io.grpc.MethodDescriptor.<com.zgg2001.grpc.service.Service.UploadInfoRequest, com.zgg2001.grpc.service.Service.UploadInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UploadParkingInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.UploadInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.UploadInfoResponse.getDefaultInstance()))
              .build();
        }
      }
    }
    return getUploadParkingInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UserLoginRequest,
      com.zgg2001.grpc.service.Service.UserLoginResponse> getUserLoginMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UserLogin",
      requestType = com.zgg2001.grpc.service.Service.UserLoginRequest.class,
      responseType = com.zgg2001.grpc.service.Service.UserLoginResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UserLoginRequest,
      com.zgg2001.grpc.service.Service.UserLoginResponse> getUserLoginMethod() {
    io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UserLoginRequest, com.zgg2001.grpc.service.Service.UserLoginResponse> getUserLoginMethod;
    if ((getUserLoginMethod = ProjectServiceGrpc.getUserLoginMethod) == null) {
      synchronized (ProjectServiceGrpc.class) {
        if ((getUserLoginMethod = ProjectServiceGrpc.getUserLoginMethod) == null) {
          ProjectServiceGrpc.getUserLoginMethod = getUserLoginMethod =
              io.grpc.MethodDescriptor.<com.zgg2001.grpc.service.Service.UserLoginRequest, com.zgg2001.grpc.service.Service.UserLoginResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UserLogin"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.UserLoginRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.UserLoginResponse.getDefaultInstance()))
              .build();
        }
      }
    }
    return getUserLoginMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UserRegistrationRequest,
      com.zgg2001.grpc.service.Service.UserRegistrationResponse> getUserRegistrationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UserRegistration",
      requestType = com.zgg2001.grpc.service.Service.UserRegistrationRequest.class,
      responseType = com.zgg2001.grpc.service.Service.UserRegistrationResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UserRegistrationRequest,
      com.zgg2001.grpc.service.Service.UserRegistrationResponse> getUserRegistrationMethod() {
    io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.UserRegistrationRequest, com.zgg2001.grpc.service.Service.UserRegistrationResponse> getUserRegistrationMethod;
    if ((getUserRegistrationMethod = ProjectServiceGrpc.getUserRegistrationMethod) == null) {
      synchronized (ProjectServiceGrpc.class) {
        if ((getUserRegistrationMethod = ProjectServiceGrpc.getUserRegistrationMethod) == null) {
          ProjectServiceGrpc.getUserRegistrationMethod = getUserRegistrationMethod =
              io.grpc.MethodDescriptor.<com.zgg2001.grpc.service.Service.UserRegistrationRequest, com.zgg2001.grpc.service.Service.UserRegistrationResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UserRegistration"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.UserRegistrationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.UserRegistrationResponse.getDefaultInstance()))
              .build();
        }
      }
    }
    return getUserRegistrationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.GetUserDataRequest,
      com.zgg2001.grpc.service.Service.GetUserDataResponse> getGetUserDataMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserData",
      requestType = com.zgg2001.grpc.service.Service.GetUserDataRequest.class,
      responseType = com.zgg2001.grpc.service.Service.GetUserDataResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.GetUserDataRequest,
      com.zgg2001.grpc.service.Service.GetUserDataResponse> getGetUserDataMethod() {
    io.grpc.MethodDescriptor<com.zgg2001.grpc.service.Service.GetUserDataRequest, com.zgg2001.grpc.service.Service.GetUserDataResponse> getGetUserDataMethod;
    if ((getGetUserDataMethod = ProjectServiceGrpc.getGetUserDataMethod) == null) {
      synchronized (ProjectServiceGrpc.class) {
        if ((getGetUserDataMethod = ProjectServiceGrpc.getGetUserDataMethod) == null) {
          ProjectServiceGrpc.getGetUserDataMethod = getGetUserDataMethod =
              io.grpc.MethodDescriptor.<com.zgg2001.grpc.service.Service.GetUserDataRequest, com.zgg2001.grpc.service.Service.GetUserDataResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserData"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.GetUserDataRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.zgg2001.grpc.service.Service.GetUserDataResponse.getDefaultInstance()))
              .build();
        }
      }
    }
    return getGetUserDataMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ProjectServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ProjectServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ProjectServiceStub>() {
        @java.lang.Override
        public ProjectServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ProjectServiceStub(channel, callOptions);
        }
      };
    return ProjectServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ProjectServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ProjectServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ProjectServiceBlockingStub>() {
        @java.lang.Override
        public ProjectServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ProjectServiceBlockingStub(channel, callOptions);
        }
      };
    return ProjectServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ProjectServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ProjectServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ProjectServiceFutureStub>() {
        @java.lang.Override
        public ProjectServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ProjectServiceFutureStub(channel, callOptions);
        }
      };
    return ProjectServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class ProjectServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void licencePlateCheck(com.zgg2001.grpc.service.Service.LPCheckRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.LPCheckResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getLicencePlateCheckMethod(), responseObserver);
    }

    /**
     */
    public void uploadParkingInfo(com.zgg2001.grpc.service.Service.UploadInfoRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UploadInfoResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getUploadParkingInfoMethod(), responseObserver);
    }

    /**
     */
    public void userLogin(com.zgg2001.grpc.service.Service.UserLoginRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UserLoginResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getUserLoginMethod(), responseObserver);
    }

    /**
     */
    public void userRegistration(com.zgg2001.grpc.service.Service.UserRegistrationRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UserRegistrationResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getUserRegistrationMethod(), responseObserver);
    }

    /**
     */
    public void getUserData(com.zgg2001.grpc.service.Service.GetUserDataRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.GetUserDataResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetUserDataMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getLicencePlateCheckMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.zgg2001.grpc.service.Service.LPCheckRequest,
                com.zgg2001.grpc.service.Service.LPCheckResponse>(
                  this, METHODID_LICENCE_PLATE_CHECK)))
          .addMethod(
            getUploadParkingInfoMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.zgg2001.grpc.service.Service.UploadInfoRequest,
                com.zgg2001.grpc.service.Service.UploadInfoResponse>(
                  this, METHODID_UPLOAD_PARKING_INFO)))
          .addMethod(
            getUserLoginMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.zgg2001.grpc.service.Service.UserLoginRequest,
                com.zgg2001.grpc.service.Service.UserLoginResponse>(
                  this, METHODID_USER_LOGIN)))
          .addMethod(
            getUserRegistrationMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.zgg2001.grpc.service.Service.UserRegistrationRequest,
                com.zgg2001.grpc.service.Service.UserRegistrationResponse>(
                  this, METHODID_USER_REGISTRATION)))
          .addMethod(
            getGetUserDataMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.zgg2001.grpc.service.Service.GetUserDataRequest,
                com.zgg2001.grpc.service.Service.GetUserDataResponse>(
                  this, METHODID_GET_USER_DATA)))
          .build();
    }
  }

  /**
   */
  public static final class ProjectServiceStub extends io.grpc.stub.AbstractAsyncStub<ProjectServiceStub> {
    private ProjectServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ProjectServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ProjectServiceStub(channel, callOptions);
    }

    /**
     */
    public void licencePlateCheck(com.zgg2001.grpc.service.Service.LPCheckRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.LPCheckResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getLicencePlateCheckMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void uploadParkingInfo(com.zgg2001.grpc.service.Service.UploadInfoRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UploadInfoResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getUploadParkingInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void userLogin(com.zgg2001.grpc.service.Service.UserLoginRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UserLoginResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getUserLoginMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void userRegistration(com.zgg2001.grpc.service.Service.UserRegistrationRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UserRegistrationResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getUserRegistrationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserData(com.zgg2001.grpc.service.Service.GetUserDataRequest request,
        io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.GetUserDataResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetUserDataMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ProjectServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<ProjectServiceBlockingStub> {
    private ProjectServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ProjectServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ProjectServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.zgg2001.grpc.service.Service.LPCheckResponse licencePlateCheck(com.zgg2001.grpc.service.Service.LPCheckRequest request) {
      return blockingUnaryCall(
          getChannel(), getLicencePlateCheckMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.zgg2001.grpc.service.Service.UploadInfoResponse uploadParkingInfo(com.zgg2001.grpc.service.Service.UploadInfoRequest request) {
      return blockingUnaryCall(
          getChannel(), getUploadParkingInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.zgg2001.grpc.service.Service.UserLoginResponse userLogin(com.zgg2001.grpc.service.Service.UserLoginRequest request) {
      return blockingUnaryCall(
          getChannel(), getUserLoginMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.zgg2001.grpc.service.Service.UserRegistrationResponse userRegistration(com.zgg2001.grpc.service.Service.UserRegistrationRequest request) {
      return blockingUnaryCall(
          getChannel(), getUserRegistrationMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.zgg2001.grpc.service.Service.GetUserDataResponse getUserData(com.zgg2001.grpc.service.Service.GetUserDataRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetUserDataMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ProjectServiceFutureStub extends io.grpc.stub.AbstractFutureStub<ProjectServiceFutureStub> {
    private ProjectServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ProjectServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ProjectServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.zgg2001.grpc.service.Service.LPCheckResponse> licencePlateCheck(
        com.zgg2001.grpc.service.Service.LPCheckRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getLicencePlateCheckMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.zgg2001.grpc.service.Service.UploadInfoResponse> uploadParkingInfo(
        com.zgg2001.grpc.service.Service.UploadInfoRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getUploadParkingInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.zgg2001.grpc.service.Service.UserLoginResponse> userLogin(
        com.zgg2001.grpc.service.Service.UserLoginRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getUserLoginMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.zgg2001.grpc.service.Service.UserRegistrationResponse> userRegistration(
        com.zgg2001.grpc.service.Service.UserRegistrationRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getUserRegistrationMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.zgg2001.grpc.service.Service.GetUserDataResponse> getUserData(
        com.zgg2001.grpc.service.Service.GetUserDataRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetUserDataMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LICENCE_PLATE_CHECK = 0;
  private static final int METHODID_UPLOAD_PARKING_INFO = 1;
  private static final int METHODID_USER_LOGIN = 2;
  private static final int METHODID_USER_REGISTRATION = 3;
  private static final int METHODID_GET_USER_DATA = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ProjectServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ProjectServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LICENCE_PLATE_CHECK:
          serviceImpl.licencePlateCheck((com.zgg2001.grpc.service.Service.LPCheckRequest) request,
              (io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.LPCheckResponse>) responseObserver);
          break;
        case METHODID_UPLOAD_PARKING_INFO:
          serviceImpl.uploadParkingInfo((com.zgg2001.grpc.service.Service.UploadInfoRequest) request,
              (io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UploadInfoResponse>) responseObserver);
          break;
        case METHODID_USER_LOGIN:
          serviceImpl.userLogin((com.zgg2001.grpc.service.Service.UserLoginRequest) request,
              (io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UserLoginResponse>) responseObserver);
          break;
        case METHODID_USER_REGISTRATION:
          serviceImpl.userRegistration((com.zgg2001.grpc.service.Service.UserRegistrationRequest) request,
              (io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.UserRegistrationResponse>) responseObserver);
          break;
        case METHODID_GET_USER_DATA:
          serviceImpl.getUserData((com.zgg2001.grpc.service.Service.GetUserDataRequest) request,
              (io.grpc.stub.StreamObserver<com.zgg2001.grpc.service.Service.GetUserDataResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (ProjectServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(getLicencePlateCheckMethod())
              .addMethod(getUploadParkingInfoMethod())
              .addMethod(getUserLoginMethod())
              .addMethod(getUserRegistrationMethod())
              .addMethod(getGetUserDataMethod())
              .build();
        }
      }
    }
    return result;
  }
}

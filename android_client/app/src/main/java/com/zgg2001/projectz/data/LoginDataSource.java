package com.zgg2001.projectz.data;

import com.zgg2001.projectz.data.model.LoggedInUser;

import com.zgg2001.grpc.service.ProjectServiceGrpc;
import com.zgg2001.grpc.service.Service;

import java.io.IOException;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

/**
 * Class that handles authentication w/ login credentials and retrieves user information.
 */
public class LoginDataSource {

    public Result<LoggedInUser> login(String username, String password) {

        try {
            // TODO: handle loggedInUser authentication
            LoggedInUser fakeUser =
                    new LoggedInUser(
                            java.util.UUID.randomUUID().toString(),
                            "Zgg2001");
            rpcLogin();
            return new Result.Success<>(fakeUser);
        } catch (Exception e) {
            return new Result.Error(new IOException("Error logging in", e));
        }
    }

    public void logout() {
        // TODO: revoke authentication
    }

    public void rpcLogin() {
        //构建通道
        final ManagedChannel channel = newChannel("xxx.xxx.xxx.xxx", 8888);
        //构建服务api代理
        ProjectServiceGrpc.ProjectServiceStub mStub = ProjectServiceGrpc.newStub(channel);
        //HelloRequest是自动生成的实体类
        Service.UserLoginRequest request = Service.UserLoginRequest.newBuilder().setUsername("zhj").setPassword("123").build();
        StreamObserver<Service.UserLoginResponse> l = new StreamObserver<Service.UserLoginResponse>() {
            @Override
            public void onNext(Service.UserLoginResponse value) {
                System.out.println(value.getResult());
            }

            @Override
            public void onError(Throwable t) {
                t.printStackTrace();
            }

            @Override
            public void onCompleted() {
                System.out.println("查询结束");
            }
        };
        mStub.userLogin(request,l); //开始请求，并拿到response
    }

    public static ManagedChannel newChannel(String host, int port) {
        return ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
    }
}
package com.zgg2001.projectz.data;

import com.zgg2001.projectz.data.model.LoggedInUser;

import com.zgg2001.grpc.service.ProjectServiceGrpc;
import com.zgg2001.grpc.service.Service;
import com.zgg2001.projectz.grpc.Operate;

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
            Operate.rpcLogin("xxx", "xxx");
            return new Result.Success<>(fakeUser);
        } catch (Exception e) {
            return new Result.Error(new IOException("Error logging in", e));
        }
    }

    public void logout() {
        // TODO: revoke authentication
    }
}
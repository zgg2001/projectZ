package com.zgg2001.projectz.grpc;

import android.net.SSLCertificateSocketFactory;

import com.zgg2001.grpc.service.ProjectServiceGrpc;
import com.zgg2001.grpc.service.Service;
import io.grpc.stub.StreamObserver;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.lang.reflect.Method;
import java.security.KeyManagementException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateException;

import java.security.KeyStore;
import java.security.cert.Certificate;
import java.security.cert.CertificateFactory;
import java.util.Arrays;

import javax.net.ssl.SSLContext;
import javax.net.ssl.SSLSocketFactory;
import javax.net.ssl.TrustManager;
import javax.net.ssl.TrustManagerFactory;
import javax.net.ssl.X509TrustManager;

import io.grpc.ManagedChannel;
import io.grpc.okhttp.OkHttpChannelBuilder;

public class Operate {
     public static void rpcLogin(String username, String password) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException, IOException {
        //构建通道
        final ManagedChannel channel = newChannel("xxx.xxx.xxx.xxx", 8888);
        //构建服务api代理
        ProjectServiceGrpc.ProjectServiceStub mStub = ProjectServiceGrpc.newStub(channel);
        //HelloRequest是自动生成的实体类
        Service.UserLoginRequest request = Service.UserLoginRequest.newBuilder().setUsername(username).setPassword(password).build();
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
        try{
            CertificateFactory cf = CertificateFactory.getInstance("X.509");
            FileInputStream inCA = new FileInputStream("android_client/app/src/main/res/auth/ca.crt");
            Certificate ca = cf.generateCertificate(inCA);
            KeyStore kStore = KeyStore.getInstance(KeyStore.getDefaultType());
            kStore.load(null, null);
            kStore.setCertificateEntry("ca", ca);
            TrustManagerFactory tmf = TrustManagerFactory
                    .getInstance(TrustManagerFactory.getDefaultAlgorithm());
            tmf.init(kStore);
            TrustManager[]  trustManagers = tmf.getTrustManagers();
            if (trustManagers.length != 1 || !(trustManagers[0] instanceof X509TrustManager)) {
                throw new IllegalStateException("Unexpected default trust managers:" + Arrays.toString(trustManagers));
            }
            SSLContext context = SSLContext.getInstance("TLS");
            context.init(null, tmf.getTrustManagers(), null);
            SSLSocketFactory sslSocketFactory = context.getSocketFactory();
            return OkHttpChannelBuilder
                    .forAddress(host, port)
                    .useTransportSecurity()
                    .overrideAuthority("IP:PORT")
                    .sslSocketFactory(sslSocketFactory)
                    .build();
        } catch (Exception e){
            e.printStackTrace();
            return null;
        }
    }
}



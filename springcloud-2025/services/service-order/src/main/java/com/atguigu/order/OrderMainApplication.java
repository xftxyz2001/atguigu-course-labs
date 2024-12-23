package com.atguigu.order;


import com.alibaba.cloud.nacos.NacosConfigManager;
import com.alibaba.cloud.nacos.configdata.NacosConfigDataLoader;
import com.alibaba.nacos.api.config.listener.Listener;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.logging.DeferredLogFactory;
import org.springframework.boot.logging.DeferredLogs;
import org.springframework.cloud.bootstrap.BootstrapConfiguration;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.context.annotation.Bean;

import java.util.concurrent.Executor;
import java.util.concurrent.Executors;


@EnableDiscoveryClient
@SpringBootApplication
public class OrderMainApplication {

    public static void main(String[] args) {
        SpringApplication.run(OrderMainApplication.class, args);
    }

    @Bean
    ApplicationRunner runner(NacosConfigManager manager){
        return args -> {
            manager.getConfigService().addListener("order.properties",
                    "DEFAULT_GROUP", new Listener() {
                        @Override
                        public Executor getExecutor() {
                            return Executors.newFixedThreadPool(4);
                        }

                        @Override
                        public void receiveConfigInfo(String configInfo) {
                            System.out.println("configInfo = " + configInfo);
                        }
                    });
    	};
    }


}

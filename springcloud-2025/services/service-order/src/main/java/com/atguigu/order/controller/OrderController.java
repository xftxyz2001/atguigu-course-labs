package com.atguigu.order.controller;


import com.atguigu.order.bean.Order;
import com.atguigu.order.service.OrderService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;


//@RefreshScope
@RestController
public class OrderController {


    @Autowired
    OrderService orderService;

    @Value("${order.timeout:777}")
    String orderTimeout;
    @Value("${order.auto-confirm:888}")
    String orderAutoConfirm;

    @GetMapping("/config")
    public String config(){

        return "超时："+orderTimeout +"; 自动确认："+orderAutoConfirm;
    }

    //创建订单
    @GetMapping("/create")
    public Order createOrder(@RequestParam("userId") Long userId,
                             @RequestParam("productId") Long productId){
        Order order = orderService.createOrder(productId, userId);
        return order;
    }
}

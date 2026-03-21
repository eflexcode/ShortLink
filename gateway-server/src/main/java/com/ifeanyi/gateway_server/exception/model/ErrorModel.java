package com.ifeanyi.gateway_server.exception.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Date;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class ErrorModel {

    private String message;
    private Date timestamp;
    private int status;
}

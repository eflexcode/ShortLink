package com.ifeanyi.url_server.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.RequiredArgsConstructor;

import java.util.Date;

@Data
@RequiredArgsConstructor
@AllArgsConstructor
public class StandardResponse {

    private String message;
    private int status;
    private Date timestamp;

}

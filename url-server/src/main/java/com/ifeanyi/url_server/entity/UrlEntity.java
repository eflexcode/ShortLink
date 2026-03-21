package com.ifeanyi.url_server.entity;

import lombok.Data;

import java.util.Date;

@Data
public class UrlEntity {

    private String id;
    private String urlOriginal;
    private String urlShort;
    private String ownerId;
    private long hits = 0;
    private Date createdAt;

}

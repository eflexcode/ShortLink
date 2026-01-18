package com.ifeanyi.url_server.entity;

import lombok.Data;
import org.springframework.data.mongodb.core.mapping.Document;

import java.math.BigInteger;
import java.util.Date;

@Data
@Document("urls")
public class UrlEntity {

    private String id;
    private String urlOriginal;
    private String urlShort;
    private String ownerId;
    private BigInteger hits;
    private Date createdAt;

}

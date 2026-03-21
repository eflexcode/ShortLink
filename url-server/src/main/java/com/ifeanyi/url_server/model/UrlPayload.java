package com.ifeanyi.url_server.model;

import lombok.Data;

@Data
public class UrlPayload {

    private String urlOriginal;
    private String ownerId;// not required as users that are not logged in can also generate shortlinks

}

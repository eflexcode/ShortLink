package com.ifeanyi.url_server.config;

import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.config.AbstractMongoClientConfiguration;

import static com.ifeanyi.url_server.Util.DB_NAME;
import static com.ifeanyi.url_server.Util.MONGO_URL;

@Configuration
public class MongoConfig extends AbstractMongoClientConfiguration {

    @Override
    protected String getDatabaseName() {
        return DB_NAME;
    }

    @Override
    public MongoClient mongoClient() {
      return MongoClients.create(MONGO_URL);
    }

}

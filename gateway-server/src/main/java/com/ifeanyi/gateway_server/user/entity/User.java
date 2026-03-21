package com.ifeanyi.gateway_server.user.entity;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import lombok.Data;
import org.jspecify.annotations.Nullable;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import java.time.LocalDateTime;
import java.util.Collection;
import java.util.List;

@Data
@Entity
public class User implements UserDetails {

    @JsonProperty("id")
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @JsonProperty("name")
    private String name;

    @JsonProperty("username")
    private String username;

    @JsonIgnore
    private String password;

    @JsonProperty("created_at")
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    private LocalDateTime updatedAt;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return List.of();
    }

    @Override
    public @Nullable String getPassword() {
        return password;
    }

    @Override
    public String getUsername() {
        return username;
    }

}

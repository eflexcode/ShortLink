package com.ifeanyi.url_server.exceptions;

import com.ifeanyi.url_server.model.StandardResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.util.Date;

@ControllerAdvice
public class GlobalException extends ResponseEntityExceptionHandler {

    @ExceptionHandler(NotFoundException.class)
    public ResponseEntity<StandardResponse> handleNotFoundExceptionHandler(NotFoundException notFoundException){
        return new ResponseEntity<StandardResponse>(new StandardResponse(notFoundException.getMessage(),HttpStatus.NOT_FOUND.value(),new Date()),HttpStatus.NOT_FOUND);
    }

}

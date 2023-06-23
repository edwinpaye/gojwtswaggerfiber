FROM scratch

# Copy binary and config files from /build 
# to root folder of scratch container.
COPY ["/build/", "/"]

# Export necessary port.
EXPOSE 3000

# Command to run when starting the container.
ENTRYPOINT ["./main"]
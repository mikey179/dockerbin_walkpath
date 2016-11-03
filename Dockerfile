FROM scratch

ADD dockerbin_walkpath /dockerbin_walkpath

ENTRYPOINT ["/dockerbin_walkpath"]

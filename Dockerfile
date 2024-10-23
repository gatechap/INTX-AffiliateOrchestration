FROM gcr.io/tdg-sec-non-prod-bnxe/go1-18-alpine as builder

WORKDIR /app/

COPY . .

#RUN go build .

RUN go get -d -v ./...

RUN go install -v ./...

#FROM alpine as production

#COPY --from=builder /app/. .

EXPOSE 8080

#CMD [.]

CMD ["affiliateorchestration"]
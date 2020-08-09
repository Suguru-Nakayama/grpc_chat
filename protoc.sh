CLIENT_OUTDIR=client/src/gen/pb
SERVER_OUTPUT_DIR=api/

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=proto ./proto/*.proto \
    --js_out=import_style=commonjs:${CLIENT_OUTDIR} \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:${CLIENT_OUTDIR} \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
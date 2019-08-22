CURPATH=$(pwd)

cd $CURPATH
mkdir -p bin/risk
cd bin/risk
go build ${CURPATH}/cmd/risk/server/
go build ${CURPATH}/cmd/risk/client/


cd $CURPATH
mkdir -p bin/db
cd bin/db
go build ${CURPATH}/cmd/db/server/
go build ${CURPATH}/cmd/db/client-grpc

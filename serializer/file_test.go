package serializer_test

import (
	"os"
	"testing"

	"github.com/adwait-godbole/learning-grpc/pb"
	"github.com/adwait-godbole/learning-grpc/sample"
	"github.com/adwait-godbole/learning-grpc/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	err := os.MkdirAll("../tmp", os.ModePerm)
	require.NoError(t, err)

	// defer os.RemoveAll("../tmp")

	laptop1 := sample.NewLaptop()
	err = serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}

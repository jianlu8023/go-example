package main

import (
	"context"

	"github.com/jianlu8023/go-tools/pkg/format/json"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/jianlu8023/go-example/pkg/logger"
)

var appLogger = logger.GetAPPLogger()

// 测试时使用go1.17.5 或许go1.17.13 可用v7.0.67
// github.com/minio/minio-go/v7 v7.0.67 // 1.17 最后的版本 测试未通过
// github.com/minio/minio-go/v7 v7.0.66 // 1.17 可用 与mongo-driver有冲突  github.com/klauspost/compress 当前使用的版本可用 replace github.com/klauspost/compress v1.17.4 => github.com/klauspost/compress v1.16.7

func main() {
	client, err := minio.New("127.0.0.1:9010", &minio.Options{
		Creds: credentials.NewStaticV4("0hEsq7JdoSihXIEgcLK2",
			"3I4DMqEE0aol3frM3M8Li51F4GyavTFkk4z0ITLc",
			""),
		Secure: false,
	})
	if err != nil {
		appLogger.Errorf("获取minio client失败: %v", err)
		return
	}
	appLogger.Infof("已经获取minio client")
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, "demo")
	if err != nil {
		appLogger.Errorf("获取bucket是否存在失败: %v", err)
		return
	}
	appLogger.Infof("bucket demo 是否存在: %v", exists)
	buckets, err := client.ListBuckets(ctx)
	if err != nil {
		appLogger.Errorf("获取bucket列表失败: %v", err)
		return
	}
	appLogger.Infof("获取bucket列表成功")

	// object, err := client.PutObject(ctx, "demo",
	// 	"demo2.txt",
	// 	strings.NewReader("this is a demo2"), int64(len("this is a demo2")),
	// 	minio.PutObjectOptions{ContentType: "application/octet-stream"})
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// fmt.Println("Successfully uploaded bytes: ", object)

	for _, bucket := range buckets {
		appLogger.Infof("当前 bucket name: %s", bucket.Name)
		objects := client.ListObjects(ctx, bucket.Name, minio.ListObjectsOptions{
			WithMetadata: true,
			WithVersions: true,
			Recursive:    true,
		})
		for obj := range objects {
			prettyJSON, err := json.PrettyJSON(obj)
			if err != nil {
				appLogger.Errorf("格式化json失败: %v", err)
				return
			}
			appLogger.Infof("bucket %s 的object信息: \n%s", bucket.Name, prettyJSON)
			// if "demo" == bucket.Name {
			// 	err := client.RemoveObject(ctx, bucket.Name, obj.Key, minio.RemoveObjectOptions{})
			// 	if err != nil {
			// 		appLogger.Errorf("删除object失败: %v", err)
			// 		return
			// 	}
			// }
		}
		// if "demo" == bucket.Name {
		// 	err := client.RemoveBucket(ctx, bucket.Name)
		// 	if err != nil {
		// 		appLogger.Errorf("删除bucket失败: %v", err)
		// 		return
		// 	}
		// }
	}
}

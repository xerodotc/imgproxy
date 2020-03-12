# Serving files from Huawei OBS

imgproxy can process images from Huawei OBS buckets. To use this feature, do the following:

1. Set `IMGPROXY_USE_OBS` environment variable as `true`;
2. [Setup credentials](#setup-credentials) to grant access to your bucket;
3. Specify Huawei OBS endpoint with `IMGPROXY_OBS_ENDPOINT`;
4. Use `obs://%bucket_name/%file_key` as the source image URL.

### Setup credentials

You can specify Huawei Cloud Access Key and Secret Key by setting the standard `IMGPROXY_OBS_ACCESSKEY` and `IMGPROXY_OBS_SECRETKEY` environment variables.

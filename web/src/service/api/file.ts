import { uploadFile } from '@tarojs/taro';
import { getRequestUrl, getRequestHeaders } from '@/service/request/helpers';

/** 文件上传 */
export async function uploadFileToSystem(tempUrl: string) {
  const config = {
    useErrMsg: false,
    contentType: "multipart/form-data",
  }
  const url = getRequestUrl("/api/file")
  const resp = await uploadFile({
    url,
    filePath: tempUrl,
    name: 'file',
    header: getRequestHeaders(config),
  });
  let fileUrl = "";
  fileUrl = JSON.parse(resp.data)?.data?.url || ""
  if (fileUrl) {
    fileUrl = getRequestUrl(`/${fileUrl}`);
  }
  console.log("fileUrl:", fileUrl);
  return fileUrl || "";
}

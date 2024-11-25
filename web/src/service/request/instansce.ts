import { request } from '@tarojs/taro';
import { REQUEST_TIMEOUT, SUCCESS_CODE, NO_AUTH_CODE } from '@/constants';
import { getRequestUrl, getRequestHeaders, showErrorMsg } from './helpers';
import { userLogin } from '@/service/api';


async function axios<T>(config: Service.RequestParam): Promise<Service.RequestResult<T>> {
  const { method, url, data } = config;
  const axiosConfig = config.axiosConfig as Service.AxiosConfig;
  const header = getRequestHeaders(axiosConfig);
  return await new Promise((resolve, reject) => {
    request({
      /** 兼容Url不同的情况，可以通过填写完整路径 */
      url: getRequestUrl(url),
      method,
      /** 对所有请求添加时间戳以防止缓存 */
      data: { _t: Date.now(), ...data },
      header,
      timeout: REQUEST_TIMEOUT,
      success: res => {
        const { code, msg, data } = res.data as Service.BackendResultConfig<T>;
        console.log("code:", code)
        /* 成功请求 */
        if (code === SUCCESS_CODE) {
          return resolve({
            error: null,
            success: data
          });
        }
        if (code === NO_AUTH_CODE) {
          userLogin();
          return resolve({
            error: null,
            success: data
          });
        }
        /** 仅有使用服务端错误信息的请求才 toast 提示错误 */
        if (axiosConfig.useErrMsg) {
          showErrorMsg(msg);
        }
        return resolve({
          error: {
            msg,
            code
          },
          success: null
        });
      },
      fail: err => {
        reject(err);
      },
      complete: () => {
        //
      }
    });
  });
}

export default axios;

import { getEnv, getAccountInfoSync, showToast } from '@tarojs/taro';
import { CONTENT_TYPE, ERROR_MSG_DURATION } from '@/constants';
import { localStg, exeStrategyActions } from '@/utils';

const env = getEnv();

/**
 * 获取请求路径
 * @param url
 */
export function getRequestUrl(url: string) {
  let baseUrl = '';
  const actions: Common.StrategyAction[] = [
    [env === 'WEB', () => (baseUrl = `/api${url}`)],
    [
      env === 'WEAPP',
      () => {
        const { miniProgram } = getAccountInfoSync();
        const hosts = {
          develop: 'http://127.0.0.1:8888', // 开发
          trial: 'https://getman.cn/mock', // 体验
          release: 'https://getman.cn/mock' // 正式
        };
        baseUrl = url.substring(0, 1) === '/' ? `${hosts[miniProgram.envVersion]}${url}` : `${url}`;
      }
    ],
    [
      true,
      () => {
        baseUrl = url.substring(0, 1) === '/' ? `${process.env.HTTP_URL}${url}` : `${url}`;
      }
    ]
  ];
  exeStrategyActions(actions);
  return baseUrl;
}

/** 获取请求头 */
export function getRequestHeaders(axiosConfig: Service.AxiosConfig) {
  const header: TaroGeneral.IAnyObject = {};
  /** 获取token */
  const token = localStg.get('token');
  if (token) {
    /** 添加token */
    header["x-token"] = token;
  }
  /** 增加类型 */
  header['Content-Type'] = axiosConfig.contentType || CONTENT_TYPE.json;
  return header;
}

export function showErrorMsg(message: string) {
  showToast({
    title: message,
    icon: 'none',
    duration: ERROR_MSG_DURATION
  });
}

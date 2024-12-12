import { getEnv, getAccountInfoSync, showToast } from '@tarojs/taro';
import { CONTENT_TYPE, ERROR_MSG_DURATION } from '@/constants';
import { localStg, exeStrategyActions } from '@/utils';
import { userLogin } from '@/service/api';

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
          trial: 'https://mp.jado.life', // 体验
          release: 'https://mp.jado.life' // 正式
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
export async function getRequestHeaders(axiosConfig: Service.AxiosConfig, needToken: boolean = true) {
  const header: TaroGeneral.IAnyObject = {};
  /** 获取token */
  let token = localStg.get('token');
  if (needToken && !token) {
    await userLogin();
    token = localStg.get('token');
  }
  /** 添加token */
  header["x-token"] = token;
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

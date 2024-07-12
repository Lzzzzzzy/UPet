import dayjs from 'dayjs';
import isoWeek from 'dayjs/plugin/isoWeek';
import type { Dayjs } from 'dayjs';
dayjs.extend(isoWeek);

/** 格式化日期时间 
 * @param date Date 对象
 * 日期格式： YYYY-MM-DD HH:mm
*/
export function formatDatetime(date: Date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hour = String(date.getHours()).padStart(2, '0');
  const minute = String(date.getMinutes()).padStart(2, '0');
  const formattedDateTime = `${year}-${month}-${day} ${hour}:${minute}`;
  return formattedDateTime
}

/** 格式化时间 
 * @param date Date 对象
 * 日期格式： HH:mm
*/
export function formatTime(date: Date) {
  const hour = String(date.getHours()).padStart(2, '0');
  const minute = String(date.getMinutes()).padStart(2, '0');
  return `${hour}:${minute}`;
}

/** 格式化日期 
 * @param date Date 对象
 * 日期格式： YYYY-MM-DD
*/
export function formatDate(date: Date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

/** 格式化日期 
 * @param date Date 对象
 * 日期格式： YYYY-MM
*/
export function formatMonth(date: Date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  return `${year}-${month}`;
}

interface DateRange {
    start: dayjs.Dayjs;
    days: number;
    end: dayjs.Dayjs;
}

export const getWeek = (index = 0, date: Date | Dayjs): DateRange => {
    const start = dayjs(date).isoWeekday(1 + 7 * index);
    const end = dayjs(date).isoWeekday(7 + 7 * index);

    return { start, days: 7, end };
};

export const getMonth = (index = 0, date: Date | Dayjs): DateRange => {
    const month = dayjs(date).add(index, 'month');
    const days = month.daysInMonth();
    
    return {
        start: dayjs(month).date(1),
        days,
        end: dayjs(month).date(days),
    };
};

export const getDays = (index = 0, date: Date | Dayjs, type: string) => {
    const { start, days } = type === 'week' ? getWeek(index, date) : getMonth(index, date);
    const list = [];

    for (let i = 0; i<days; i++) {
      const currentDate = start.add(i, "day");
      list.push({
        d: currentDate.date(),
        m: currentDate.month() + 1,
        date: currentDate.format('YYYY-MM-DD'),
        day: currentDate,
      });
    }

    return list;
};

/** 判断date是否在date列表中 
 * @param dateToCheck Date 对象
 * @param dateArray Date 列表
*/
export const isDateInArray = (dateToCheck: Date, dateArray: Array<Date>) => {
  // 将dateToCheck转换为只有年月日的日期
  const checkDateOnly = new Date(dateToCheck.getFullYear(), dateToCheck.getMonth(), dateToCheck.getDate());

  // 遍历dateArray并检查每个日期
  for (let i = 0; i < dateArray.length; i++) {
    // 将数组中的每个日期也转换为只有年月日的日期
    const arrayDateOnly = new Date(dateArray[i].getFullYear(), dateArray[i].getMonth(), dateArray[i].getDate());

    // 如果两个日期相同（仅比较年月日），则返回true
    if (checkDateOnly.getTime() === arrayDateOnly.getTime()) {
      return true;
    }
  }

  // 如果没有找到匹配的日期，则返回false
  return false;
}

/** 判断两个日期是否是同一天
 * @param date1
 * @param date2
 */
export const isSameDate = (date1: Date | string | Dayjs, date2: Date | string | Dayjs) => {
  return dayjs(date1).format("YYYY-MM-DD") === dayjs(date2).format("YYYY-MM-DD");
}

/** 判断日期1在日期2之前
 * @param date1
 * @param date2
 */
export const isBeforeDate = (date1: Date | string | Dayjs, date2: Date | string | Dayjs) => {
  return dayjs(dayjs(date1).format("YYYY-MM-DD")).isBefore(dayjs(date2).format("YYYY-MM-DD"));
}

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

/** 获取可选的当天时间
 * @param date Date 对象
 */
export function getMinAndMaxDate(date: Date) {
  const today = new Date();
  let minDate;
  const maxDate = new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59);

  if (date.getMonth() === today.getMonth() && date.getFullYear() === today.getFullYear() && date.getDate() === today.getDate()) {
    minDate = new Date(date.getFullYear(), date.getMonth(), date.getDate(), date.getHours(), date.getMinutes(), 0);
  } else {
    minDate = new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0);
  }
  return {minDate, maxDate}
}


/** 格式化日期 
 * @param date Date 对象
 * 日期格式： YYYY-MM-DD
*/
export function formatDate(date: Date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const formattedDate = `${year}-${month}-${day}`;
  return formattedDate
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

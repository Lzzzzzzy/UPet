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
    const days = start.daysInMonth();
    const end = dayjs(date).isoWeekday(7 + 7 * index);

    return { start, days, end };
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
    const { start, days, end } = type === 'week' ? getWeek(index, date) : getMonth(index, date);
    const list = [];
    
    const startD = start.date();
    const startM = start.month();
    const endD = end.date();
    const endM = end.month();
    
    const currentDateList = [];
    const formattedDateList = [];
    
    for (let i = startD; i <= days; i++) {
        const currentDate = start.date(i);
        const formattedDate = currentDate.format('YYYY-MM-DD');
        
        currentDateList.push(currentDate);
        formattedDateList.push(formattedDate);
        
        if (startM !== endM && i >= endD) {
            break;
        }
    }
    
    for (let i = 0; i < currentDateList.length; i++) {
        const currentDate = currentDateList[i];
        const formattedDate = formattedDateList[i];
        
        list.push({
            d: currentDate.date(),
            m: currentDate.month() + 1,
            date: formattedDate,
            day: currentDate,
        });
    }

    return list;
};

// export const getWeekDays = (index = 0, date: Date) => {
//     const {
//         start,
//         start: { $D: startD, $M: startM },
//         days,
//         end,
//         end: { $D: endD, $M: endM }
//     } = getWeek(index, date);
//     const list = [];
//     const endFlag = startM === endM ? endD : days;
//     for (let i = startD; i <= endFlag; i++) {
//         let date = start.date(i).format('YYYY-MM-DD')
//         list.push({
//             d: i,
//             m: startM + 1,
//             date,
//         });
//     }
//     if (startM !== endM) {
//         for (let i = 1; i <= endD; i++) {
//             let date = end.date(i).format('YYYY-MM-DD')
//             list.push({
//                 d: i,
//                 m: endM + 1,
//                 date,
//             });
//         }
//     }
//     return list;
// };


// export const getMonth = (index = 0, date: Date) => {
//     const month = dayjs(date).add(index, 'month');
//     const days = Array.from({ length: month.daysInMonth() }, (_, i) => {
//         return dayjs(month)
//             .date(1 + i)
//             .format('YYYY-MM-DD');
//     });
//     return {
//         start: dayjs(month).date(1).day(),
//         days,
//         month: dayjs(month).format('YYYY-MM')
//     };
// };
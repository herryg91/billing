const getWeekNum = (date: Date) => {
  const janFirst = new Date(date.getFullYear(), 0, 1);
  // Source: https://stackoverflow.com/a/27125580/3307678
  return Math.ceil((((date.getTime() - janFirst.getTime()) / 86400000) + janFirst.getDay() + 1) / 7);
}


  
export const isSameWeek = (dateA:Date, dateB:Date) => {
  console.log( (dateA.getFullYear() *100) + getWeekNum(dateA) , (dateB.getFullYear() *100) + getWeekNum(dateB))
  return (dateA.getFullYear() *100) + getWeekNum(dateA) === (dateB.getFullYear() *100) + getWeekNum(dateB);
}
  
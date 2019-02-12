// Input: [‘Oct 7, 2009’, ‘Nov 10, 2009’, ‘Jan 10, 2009’, ‘Oct 22, 2009’, …]

// Output = ['Dec 1, 2019', 'Sep 20, 2010', 'Nov 10, 2009', 'Oct 22, 2009', 'Oct 7, 2009', Jan 10, 2009']

const dateArr = ['Oct 7, 2009', 'Nov 10, 2009', 'Jan 10, 2009', 'Oct 22, 2009', 'Dec 1, 2019', 'Sep 20, 2010', 'Aug 2, 1912']

const dateSorting = (arr) => {
  const monthArr = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
  let result = []
  let tmp
  let dateObjArr = dateArrToObj(arr, monthArr)
  let sorted = sortDates(dateObjArr)

  // loop through sorted array and format w/ string literal - push to result array
  for (let i = 0; i < sorted.length; i++) {
    tmp = `${monthConversionToString(sorted[i].month, monthArr)} ${sorted[i].day}, ${sorted[i].year}`
    result.push(tmp)
  }
  return result
}

// converts date string to object {month, day, year}
const dateArrToObj = (arr, monthArr) => {
  let dateObjArr = []
  let day = 0
  for (let i = 0; i < arr.length; i++) {
    dayArr = arr[i].split(' ')[1].split('')
    for (let j = 0; j < dayArr.length; j++) {
      if (dayArr[1] == ',') {
        day = dayArr[0]
      } else {
        day = dayArr[0] + dayArr[1]
      }
    }
    dateObjArr.push({
      'month': monthConversionToNumber(arr[i].split(' ')[0], monthArr),
      'day': +day,
      'year': +(arr[i].split(' ')[2])
    })
  }
  return dateObjArr
}

// sort the date objects by key value
const sortDates = (arr) => {
  let ordered = []
  ordered = arr.sort((a, b) => {
    if (a.year === b.year) {
      if (a.month === b.month) {
        return a.day < b.day
      } return a.month < b.month
    } else {
      return a.year < b.year
    }
  })
  return ordered
}

// converts month string to integer value
const monthConversionToNumber = (str, monthArr) => {
  let num = 0
  for (let i = 0; i < monthArr.length; i++) {
    if (str === monthArr[i])
      num = i
  }
  return num
}

// converts the integer value back to month string
const monthConversionToString = (num, monthArr) => {
  let str = ''
  for (let i = 0; i < monthArr.length; i++) {
    if (num === i)
      str = monthArr[i]
  }
  return str
}

// call the f(x) with input array
dateSorting(dateArr)
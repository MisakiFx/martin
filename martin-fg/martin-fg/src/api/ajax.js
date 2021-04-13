import axios from "axios";

export default function ajax(url = '', params = {}, headers = {}, type = 'GET') {
  //1.定义promise
  let promise;
  return new Promise((resolve, reject)=>{
    if('GET' === type) {
      let paramsStr = '';
      Object.keys(params).forEach(key=>{
        paramsStr += key + '=' + params[key] + '&'
      })

      if(paramsStr !== '') {
        paramsStr = paramsStr.substr(0, paramsStr.lastIndexOf('&'))
      }

      url += '?' + paramsStr

      if (Object.keys(headers).length === 0) {
        promise = axios.get(url)
      } else {
        promise = axios.get(url, {
          headers: headers
        })
      }
    } else if('POST' === type) {
      if (Object.keys(headers).length === 0) {
        promise = axios.post(url, params)
      } else {
        promise = axios.post(url, params, {
          headers: headers
        })
      }
    }
    promise.then((response)=>{
      resolve(response.data)
    }).catch(error=>{
      reject(error)
    })
  })
}

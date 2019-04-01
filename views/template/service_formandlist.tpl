save${componentname$}(formdata,listdata): Observable<any> {
    
    return this.http.post(this.ls.api_url + '/${componentname$}/save${componentname$}', {      
      '${componentname$}': formdata,
      '${componentname$}item': listdata
    }, httpOptions).pipe();
  }

get${componentname$}(queryitem): Observable<any> {
    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}', queryitem, httpOptions).pipe();
}

get${componentname$}item(queryitem): Observable<any> {
    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}item', queryitem, httpOptions).pipe();
}
get${componentname$}byid(queryitem): Observable<any> {
    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}byid', queryitem, httpOptions).pipe();
}

get${componentname$}filesbyid(queryitem): Observable<any> {
    return this.http.post(this.ls.api_url + '/common/getfilelistbyfiid', queryitem, httpOptions).pipe();
}
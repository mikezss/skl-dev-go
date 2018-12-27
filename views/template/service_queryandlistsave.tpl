get${componentname$}count(queryitems): Observable<any> {

    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}count', queryitems, httpOptions).pipe();
}

get${componentname$}bypageindex(queryitems): Observable<any> {

    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}bypageindex', queryitems, httpOptions).pipe();
}
get${componentname$}(queryitems): Observable<any> {

    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}', queryitems,httpOptions).pipe();
}
do${componentname$}(data): Observable<any> {
    return this.http.post(this.ls.api_url + '/${componentname$}/do${componentname$}', data, httpOptions).pipe();
}

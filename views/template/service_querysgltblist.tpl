save${componentname$}(data: any): Observable<any> {

    return this.http.post(this.ls.api_url + '/${componentname$}/save${componentname$}', data, httpOptions).pipe();
}

get${componentname$}(): Observable<any> {

    return this.http.get(this.ls.api_url + '/${componentname$}/get${componentname$}', httpOptions).pipe();
}
get${componentname$}count(queryitems): Observable<any> {

    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}count', queryitems, httpOptions).pipe();
}

get${componentname$}bypageindex(queryitems): Observable<any> {

    return this.http.post(this.ls.api_url + '/${componentname$}/get${componentname$}bypageindex', queryitems, httpOptions).pipe();
}

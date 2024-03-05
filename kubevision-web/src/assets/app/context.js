
export class Context{
    constructor() {
    }
    setNamaespace(region){
        return localStorage.setItem('namespace', region)
    }
    getNamespace(){
        return localStorage.getItem('region')
    }
}

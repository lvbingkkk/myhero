function myHref(href,pro){

    return window.location.href = window.location.protocol + "//" + href + "/createHero?pro="+pro
    // return window.location.href=href + "createHero?pro="+pro
}

function myColor(r,g,b,a){
    return "background-color:rgba"+(r, g, b, a)


}

export{
    myColor,myHref
}

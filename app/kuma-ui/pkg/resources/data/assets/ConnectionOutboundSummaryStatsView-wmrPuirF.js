import{d as C,a as n,o as t,b as c,w as o,e as s,m as p,f as l,E as x,A as y,a4 as R,q as d,K as k}from"./index-FFDA7vTv.js";import{C as w}from"./CodeBlock-dbnfu_z5.js";const $=C({__name:"ConnectionOutboundSummaryStatsView",setup(S){return(v,E)=>{const m=n("RouteTitle"),_=n("KButton"),u=n("DataSource"),f=n("AppView"),h=n("RouteView");return t(),c(h,{params:{codeSearch:"",codeFilter:!1,codeRegExp:!1,mesh:"",dataPlane:"",service:""},name:"connection-outbound-summary-stats-view"},{default:o(({route:e})=>[s(f,null,{title:o(()=>[p("h3",null,[s(m,{title:"Stats"})])]),default:o(()=>[l(),p("div",null,[s(u,{src:`/meshes/${e.params.mesh}/dataplanes/${e.params.dataPlane}/stats/${e.params.service}`},{default:o(({data:r,error:i,refresh:g})=>[i?(t(),c(x,{key:0,error:i},null,8,["error"])):r===void 0?(t(),c(y,{key:1})):(t(),c(w,{key:2,language:"json",code:r.raw.split(`
`).filter(a=>a.includes(`.${e.params.service}.`)).map(a=>a.replace(`${e.params.service}.`,"")).join(`
`),"is-searchable":"",query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:a=>e.update({codeSearch:a}),onFilterModeChange:a=>e.update({codeFilter:a}),onRegExpModeChange:a=>e.update({codeRegExp:a})},{"primary-actions":o(()=>[s(_,{appearance:"primary",onClick:g},{default:o(()=>[s(d(R),{size:d(k)},null,8,["size"]),l(`
                Refresh
              `)]),_:2},1032,["onClick"])]),_:2},1032,["code","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"]))]),_:2},1032,["src"])])]),_:2},1024)]),_:1})}}});export{$ as default};
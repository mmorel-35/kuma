import{f}from"./kong-icons.es366-QijJImZw.js";import{C as h}from"./CodeBlock-CT-0xSyp.js";import{d as C,r as n,o as g,m as x,w as t,b as o,e as c,l as R}from"./index-BH25Q5I8.js";const F=C({__name:"ConnectionOutboundSummaryClustersView",setup(w){return(y,V)=>{const r=n("RouteTitle"),p=n("KButton"),i=n("DataCollection"),d=n("DataLoader"),l=n("AppView"),m=n("RouteView");return g(),x(m,{params:{codeSearch:"",codeFilter:!1,codeRegExp:!1,mesh:"",dataPlane:"",connection:""},name:"connection-outbound-summary-clusters-view"},{default:t(({route:e})=>[o(r,{render:!1,title:"Clusters"}),c(),o(l,null,{default:t(()=>[o(d,{src:`/meshes/${e.params.mesh}/dataplanes/${e.params.dataPlane}/data-path/clusters`},{default:t(({data:u,refresh:_})=>[o(i,{items:u.split(`
`),predicate:s=>s.startsWith(`${e.params.connection}::`)},{default:t(({items:s})=>[o(h,{language:"json",code:s.map(a=>a.replace(`${e.params.connection}::`,"")).join(`
`),"is-searchable":"",query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:a=>e.update({codeSearch:a}),onFilterModeChange:a=>e.update({codeFilter:a}),onRegExpModeChange:a=>e.update({codeRegExp:a})},{"primary-actions":t(()=>[o(p,{appearance:"primary",onClick:_},{default:t(()=>[o(R(f)),c(`

                Refresh
              `)]),_:2},1032,["onClick"])]),_:2},1032,["code","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]),_:2},1032,["items","predicate"])]),_:2},1032,["src"])]),_:2},1024)]),_:1})}}});export{F as default};
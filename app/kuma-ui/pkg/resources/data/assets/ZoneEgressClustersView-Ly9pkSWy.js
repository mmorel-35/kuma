import{_ as l}from"./EnvoyData.vue_vue_type_script_setup_true_lang-DD2XdIiN.js";import{d,r as s,o as m,m as u,w as t,b as a,e as _}from"./index-BH25Q5I8.js";import"./kong-icons.es366-QijJImZw.js";import"./CodeBlock-CT-0xSyp.js";const E=d({__name:"ZoneEgressClustersView",setup(g){return(f,h)=>{const n=s("RouteTitle"),r=s("KCard"),c=s("AppView"),p=s("RouteView");return m(),u(p,{name:"zone-egress-clusters-view",params:{zoneEgress:"",codeSearch:"",codeFilter:!1,codeRegExp:!1}},{default:t(({route:e,t:i})=>[a(n,{render:!1,title:i("zone-egresses.routes.item.navigation.zone-egress-clusters-view")},null,8,["title"]),_(),a(c,null,{default:t(()=>[a(r,null,{default:t(()=>[a(l,{resource:"Zone",src:`/zone-egresses/${e.params.zoneEgress}/data-path/clusters`,query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter,"is-reg-exp-mode":e.params.codeRegExp,onQueryChange:o=>e.update({codeSearch:o}),onFilterModeChange:o=>e.update({codeFilter:o}),onRegExpModeChange:o=>e.update({codeRegExp:o})},null,8,["src","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{E as default};
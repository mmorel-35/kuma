import{v as d}from"./kongponents.es-186331d0.js";import{d as p,a as m,x as o,o as n,b as r,z as v,q as i,w as a,g as l,h as w,t as f,s as x,e as _}from"./index-f9d036ef.js";import{_ as y}from"./_plugin-vue_export-helper-c27b6911.js";const V=p({__name:"MeshView",setup(g){const h=m(),t=[{hash:"mesh-detail-view",title:"Overview"},{hash:"services-list-view",title:"Services"},{hash:"gateways-list-view",title:"Gateways"},{hash:"data-planes-list-view",title:"Data Plane Proxies"},{hash:"policies",title:"Policies"}];return(k,C)=>{const c=o("router-link"),u=o("RouterView");return n(),r(_(d),{class:"route-mesh-view-tabs",tabs:t,"model-value":(t.find(e=>{var s;return(((s=_(h).currentRoute)==null?void 0:s.value.name)??"").toString().startsWith(e.hash)})??t[0]).hash},v({_:2},[i(t,e=>({name:`${e.hash}-anchor`,fn:a(()=>[l(c,{to:{name:e.hash}},{default:a(()=>[w(f(e.title),1)]),_:2},1032,["to"])])})),i(t,e=>({name:e.hash,fn:a(()=>[l(u,null,{default:a(s=>[(n(),r(x(s.Component),{key:s.route.path}))]),_:1})])}))]),1032,["model-value"])}}});const b=y(V,[["__scopeId","data-v-82dd2d41"]]);export{b as default};

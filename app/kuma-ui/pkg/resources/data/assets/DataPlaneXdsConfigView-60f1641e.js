import{d as r,o as d,a as i,w as t,h as e,i as l,b as s,g as p,k as m}from"./index-3b0c93cd.js";import{e as c,A as u,_ as f}from"./RouteView.vue_vue_type_script_setup_true_lang-eb8bc58a.js";import{_}from"./RouteTitle.vue_vue_type_script_setup_true_lang-275bb3bf.js";import{E as x}from"./EnvoyData-79e4183d.js";import{a as g}from"./dataplane-30467516.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-bd9f7fe8.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-73d87df4.js";const C=r({__name:"DataPlaneXdsConfigView",props:{data:{}},setup(n){const a=n,{t:o}=c();return(h,w)=>(d(),i(f,{name:"data-plane-xds-config-view","data-testid":"data-plane-xds-config-view"},{default:t(()=>[e(u,null,{title:t(()=>[l("h2",null,[e(_,{title:s(o)("data-planes.routes.item.navigation.data-plane-xds-config-view"),render:!0},null,8,["title"])])]),default:t(()=>[p(),e(s(m),null,{body:t(()=>[e(x,{status:s(g)(a.data.dataplane,a.data.dataplaneInsight).status,resource:"Data Plane Proxy",src:`/meshes/${a.data.mesh}/dataplanes/${a.data.name}/data-path/xds`,"query-key":"envoy-data-xds-data-plane"},null,8,["status","src"])]),_:1})]),_:1})]),_:1}))}});export{C as default};
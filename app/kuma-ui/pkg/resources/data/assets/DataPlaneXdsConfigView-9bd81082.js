import{E as l}from"./EnvoyData-54e7d3f5.js";import{d as c,a as t,o as m,b as u,w as o,e as n,p as _,f as g}from"./index-19354c6c.js";import"./index-fce48c05.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-a4faa3a1.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-ac8206e3.js";import"./ErrorBlock-f71f64d1.js";import"./TextWithCopyButton-1c7354cf.js";import"./CopyButton-2ee425f1.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-706de663.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-3651c4e8.js";const B=c({__name:"DataPlaneXdsConfigView",setup(f){return(h,x)=>{const r=t("RouteTitle"),p=t("KCard"),s=t("AppView"),d=t("RouteView");return m(),u(d,{name:"data-plane-xds-config-view",params:{mesh:"",dataPlane:"",codeSearch:"",codeFilter:!1,codeRegExp:!1}},{default:o(({route:e,t:i})=>[n(s,null,{title:o(()=>[_("h2",null,[n(r,{title:i("data-planes.routes.item.navigation.data-plane-xds-config-view")},null,8,["title"])])]),default:o(()=>[g(),n(p,null,{default:o(()=>[n(l,{resource:"Data Plane Proxy",src:`/meshes/${e.params.mesh}/dataplanes/${e.params.dataPlane}/data-path/xds`,query:e.params.codeSearch,"is-filter-mode":e.params.codeFilter==="true","is-reg-exp-mode":e.params.codeRegExp==="true",onQueryChange:a=>e.update({codeSearch:a}),onFilterModeChange:a=>e.update({codeFilter:a}),onRegExpModeChange:a=>e.update({codeRegExp:a})},null,8,["src","query","is-filter-mode","is-reg-exp-mode","onQueryChange","onFilterModeChange","onRegExpModeChange"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{B as default};
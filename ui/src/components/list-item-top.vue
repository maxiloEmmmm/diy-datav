<template>
    <div class="ysz ysz-list-item-top" @click="e => $emit('click', e)" :class="_class">
        <div class="ysz-list-item-top__top-item"><slot name="top"></slot></div>
        <div class="ysz-list-item-top__bottom-item"><slot></slot></div>
    </div>
</template>

<script>
export default {
    name: 'yszListItemTop',
    props: {
        noTm: {
            type: Boolean,
            default: false
        },
        noP: {
            type: Boolean,
            default: false
        },
        topItemCenter: {
            type: Boolean,
            default: false
        },
        bottomItemCenter: {
            type: Boolean,
            default: false
        },
        topItemEnd: {
            type: Boolean,
            default: false
        },
        bottomItemEnd: {
            type: Boolean,
            default: false
        },
        itemCenter: {
            type: Boolean,
            default: false
        },
        itemEnd: {
            type: Boolean,
            default: false
        }
    },
    computed: {
        _class(){
            return {
                'ysz-list-item-top__no-p': this.noP,
                'ysz-list-item-top__no-tm': this.noTm,
                'ysz-list-item-top__top-item-center': this.topItemCenter,
                'ysz-list-item-top__top-item-end': this.topItemEnd,
                'ysz-list-item-top__bottom-item-center': this.bottomItemCenter,
                'ysz-list-item-top__bottom-item-end': this.bottomItemEnd,
                'ysz-list-item-top__item-center': this.itemCenter,
            }
        }
    }
}
</script>

<style lang="scss">
.ysz-list-item-top {
    display: flex;
    flex-direction: column;
    padding: 2px 4px;
    flex-grow: 1;
&.ysz-list-item-top__no-p {
     padding: 0;
 }
& > .ysz-list-item-top__top-item, & > .ysz-list-item-top__bottom-item {
    display: flex;
    align-items: center;
    justify-content: flex-start;
// if top or bottom nothing will like row.
flex-grow: 1;
}
& > .ysz-list-item-top__top-item {
      margin-bottom: 0.5em;
  }
&.ysz-list-item-top__no-tm > .ysz-list-item-top__top-item {
     margin-bottom:0;
 }
&.ysz-list-item-top__top-item-center > .ysz-list-item-top__top-item,
&.ysz-list-item-top__bottom-item-center > .ysz-list-item-top__bottom-item,
&.ysz-list-item-top__item-center > .ysz-list-item-top__top-item,
&.ysz-list-item-top__item-center > .ysz-list-item-top__bottom-item{
     justify-content: center;
 }
&.ysz-list-item-top__top-item-end > .ysz-list-item-top__top-item,
&.ysz-list-item-top__bottom-item-end > .ysz-list-item-top__bottom-item,
&.ysz-list-item-top__item-end > .ysz-list-item-top__top-item,
&.ysz-list-item-top__item-end > .ysz-list-item-top__bottom-item {
     justify-content: flex-end;
 }
}
</style>
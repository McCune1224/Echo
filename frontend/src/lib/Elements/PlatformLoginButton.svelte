
<script lang='ts'>  
    import Fa from 'svelte-fa' 

    export var text;
    export var color = "bg-white";
    export var icon;
    export var postUrl;

    const OnClick = () =>
    {
        window.open(postUrl, "_self");
    }
</script>

<div class='px-1 full-size'> 
    <div class={"h-[65px] rounded-2xl w-[65px] font-[16px]  flex flex-col items-center justify-center "} on:click={OnClick}>
        <Fa class={'transition-fast w-3/4 h-3/4 text-9xl drop-shadow-shine-grey hover:drop-shadow-shine-white ' + color} icon={icon} />{text}
    </div> 
</div>
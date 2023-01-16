<script lang="ts">

    var state = false;

    function OnClick()
    {
        state = !state;
        console.log(state);
    }


</script>

<div> 

</div>

<div>
    
</div>

<button type="button" on:click={OnClick}>My Button</button>

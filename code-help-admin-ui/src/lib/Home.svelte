<script>
  import { Icon } from "svelte-icons-pack";
  import { AiFillQuestionCircle, AiOutlineBarChart } from "svelte-icons-pack/ai";
  import { RiCommunicationChat1Fill } from "svelte-icons-pack/ri";
  import { link } from "svelte-spa-router";
  import { Route } from "../routes";
  import { BiCategory } from "svelte-icons-pack/bi";
</script>

<style>
  .tile {
    width: 200px;
    height: 120px;
    background-color: #18283b;
    color: #ffffff;
    padding: 1rem;
    border-radius: 0.625rem;
    display: flex;
    gap: 0.5rem;
    text-decoration: none;
  }

  .tile:hover {
    opacity: 0.8;
  }
  .page-title {
    padding-bottom: 5rem;
    font-size: 2em;
    font-weight: bold;
  }

  .page-content {
    padding-inline: 2rem;
    padding-block: 1.5rem;
  }
</style>

<div class="p-1 column h-100 gap-1 page-content">
  <h1 class="page-title">Code Help Admin</h1>
  <h2>Core</h2>
  <hr />
  <div class="row gap-1">
    <a class="tile" href={Route.categories_overview} use:link>
      <Icon src={BiCategory} size="32" />
      <span>Categories</span>
    </a>
    <a class="tile" href={Route.problems_overview} use:link>
      <Icon src={AiFillQuestionCircle} size="32" />
      <span>Problems</span>
    </a>
    <a class="tile" href={Route.contests_overview} use:link>
      <Icon src={AiOutlineBarChart} size="32" />
      <span>Contests</span>
    </a>
  </div>
  <h2>Forum</h2>
  <hr />
  <div class="row gap-1">
    <a class="tile" href={Route.forum_categories_overview} use:link>
      <Icon src={BiCategory} size="32" />
      <span>Categories</span>
    </a><a class="tile" href={Route.communities_overview} use:link>
      <Icon src={RiCommunicationChat1Fill} size="32" />
      <span>Communities</span>
    </a>
  </div>
</div>

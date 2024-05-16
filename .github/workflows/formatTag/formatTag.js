export function formatTag(tag) {
  const parsed = /(?:^@.*\/(?<package>.*)@v?)?(?<semver>\d+.\d+.\d+)-?.*/.exec(
    tag
  );
  const parsedPackage = /(?<package>.*)@v?-?.*/.exec(tag);

  if (parsed?.groups) {
    const isSubPackage = typeof parsed.groups.package === "string";
    const pkg = isSubPackage
      ? parsed.groups.package
      : parsedPackage?.groups?.package ?? "wapi.js";
    const semver = parsed.groups.semver;

    return {
      isSubpackage: isSubPackage,
      package: pkg,
      semver,
    };
  }

  return null;
}

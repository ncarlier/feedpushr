export default (...names: (string | undefined | null)[]) => names.filter(name => name).join(' ')

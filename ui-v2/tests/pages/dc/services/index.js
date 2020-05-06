export default function(visitable, clickable, text, attribute, collection, filter) {
  const service = {
    name: text('a span:nth-child(2)'),
    service: clickable('a'),
    externalSource: attribute('data-test-external-source', '[data-test-external-source]'),
    kind: attribute('data-test-kind', '[data-test-kind]'),
  };
  return {
    visit: visitable('/:dc/services'),
    services: collection('.consul-service-list > ul > li:not(:first-child)', service),
    dcs: collection('[data-test-datacenter-picker]', {
      name: clickable('a'),
    }),
    filter: filter,
    home: clickable('[data-test-home]'),
  };
}

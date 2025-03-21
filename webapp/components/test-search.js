/**
 * Copyright 2018 The WPT Dashboard Project. All rights reserved.
 * Use of this source code is governed by a BSD-style license that can be
 * found in the LICENSE file.
 */

import '../node_modules/@polymer/paper-tooltip/paper-tooltip.js';
import { html } from '../node_modules/@polymer/polymer/polymer-element.js';
import { PolymerElement } from '../node_modules/@polymer/polymer/polymer-element.js';
import { WPTFlags } from './wpt-flags.js';
import './ohm.js';
import { AllBrowserNames } from './product-info.js';

const statuses = [
  'pass',
  'ok',
  'error',
  'timeout',
  'notrun',
  'fail',
  'crash',
  'skip',
  'assert',
  'unknown',
  'missing', // UI calls unknown missing.
];

const atoms = {
  status: statuses,
};

for (const b of AllBrowserNames) {
  atoms[b] = statuses;
}

/* global ohm */
const QUERY_GRAMMAR = ohm.grammar(`
  Query {
    Root = ListOf<OrQ, space*>

    OrQ = NonemptyListOf<AndQ, or>

    AndQ = NonemptyListOf<Q, and>

    Q = All
      | None
      | Count
      | Sequential
      | Exists

    All = "all(" ListOf<Exp, space*> ")"

    None = "none(" ListOf<Exp, space*> ")"

    Sequential = "seq(" ListOf<Exp, space*> ")"

    Count = CountSpecifier "(" Exp ")"

    CountSpecifier
      = "count" ":"? inequality number -- countInequality
      | "count:" number                -- countN
      | "three"                        -- count3
      | "two"                          -- count2
      | "one"                          -- count1

    Exists
      = "exists(" ListOf<Exp, space*> ")" -- explicit
      | AndPart                           -- implicit


    Exp = NonemptyListOf<OrPart, or>

    OrPart = NonemptyListOf<AndPart, and>

    AndPart
      = NestedExp
      | Fragment    -- fragment

    NestedExp
      = "(" Exp ")"   -- paren
      | not NestedExp -- not

    or
      = "|"
      | caseInsensitive<"or">

    and
      = "&"
      | caseInsensitive<"and">

    not
      = "!"
      | "not"

    inequality
      = ">="
      | "<="
      | ">"
      | "<"
      | "="

    Fragment
      = not Fragment -- not
      | linkExp
      | isExp
      | triagedExp
      | labelExp
      | webFeatureExp
      | statusExp
      | subtestExp
      | pathExp
      | patternExp

    statusExp
      = caseInsensitive<"status"> ":" statusLiteral  -- eq
      | caseInsensitive<"status"> ":!" statusLiteral -- neq
      | productSpec ":" statusLiteral                -- product_eq
      | productSpec ":!" statusLiteral               -- product_neq

    subtestExp
      = caseInsensitive<"subtest"> ":" nameFragment

    pathExp
      = caseInsensitive<"path"> ":" nameFragment

    linkExp
      = caseInsensitive<"link"> ":" nameFragment

    triagedExp
      = caseInsensitive<"triaged"> ":" browserName
      | caseInsensitive<"triaged"> ":" "test-issue"

    labelExp
      = caseInsensitive<"label"> ":" nameFragment

    webFeatureExp
      = caseInsensitive<"feature"> ":" nameFragment

    isExp
      = caseInsensitive<"is"> ":" metadataQualityLiteral

    patternExp = nameFragment

    productSpec = browserName ("-" browserVersion)?

    browserName
      = ${AllBrowserNames.map(b => 'caseInsensitive<"' + b + '">').join('\n      |')}

    browserVersion = number ("." number)*

    statusLiteral
      = ${statuses.map(s => 'caseInsensitive<"' + s + '">').join('\n      |')}

    metadataQualityLiteral
      = caseInsensitive<"different">
      | caseInsensitive<"tentative">
      | caseInsensitive<"optional">

    nameFragment
      = basicNameFragment                       -- basic
      | quotemark complexNameFragment quotemark -- quoted

    basicNameFragment = basicNameFragmentChar+

    complexNameFragment = nameFragmentChar+ (space+ nameFragmentChar+)*

    basicNameFragmentChar
      = letter
      | digit
      | "/"
      | "."
      | "-"
      | "_"
      | "?"

    nameFragmentChar
      = "\\x00".."\\x08"
      | "\\x0E".."\\x1F"
      | "\\x21"
      | "\\x23".."\\uFFFF"

    number = digit+
    quotemark = "\\""
    backslash = "\\\\"
  }
`);
/* eslint-disable */
const evalNot = (n, p) => {
  return {not: p.eval()};
};
const evalSelf = p => p.eval();
const emptyQuery = Object.freeze({exists: [{pattern: ''}]});
const andConjunction = l => {
  const ps = l.eval();
  return ps.length === 1 ? ps[0] : {and: ps};
};
const orConjunction = l => {
  const ps = l.eval();
  return ps.length === 1 ? ps[0] : {or: ps};
};
const QUERY_SEMANTICS = QUERY_GRAMMAR.createSemantics().addOperation('eval', {
  _terminal: function() {
    return this.sourceString;
  },
  Root: (r) => {
    const ps = r.eval();
    if (ps.length === 0) {
      return emptyQuery;
    }
    // If there's only separate implicit exists at the root, collapse them.
    const isImplicitExists = p => 'exists' in p && p.exists.length === 1
        || 'and' in p && p.and.every(isImplicitExists)
        || 'or' in p && p.or.every(isImplicitExists);
    if (ps.every(isImplicitExists)) {
      const unwrap = p => 'exists' in p && p.exists[0]
        || 'or' in p && { or: p.or.map(unwrap) }
        || 'and' in p && { and: p.and.map(unwrap) }
        || p;
      return { exists: ps.map(unwrap) };
    }
    if (ps.length === 1) {
      return ps[0];
    }
    return { and: ps };
  },
  OrQ: orConjunction,
  AndQ: andConjunction,
  EmptyListOf: function() {
    return [];
  },
  NonemptyListOf: function(fst, seps, rest) {
    return [fst.eval()].concat(rest.eval());
  },
  Exists_explicit: (l, e, r) => {
    return { exists: e.eval() };
  },
  Exists_implicit: e => {
    return { exists: [e.eval()] };
  },
  All: (_, l, __) => {
    const ps = l.eval();
    return ps.length === 0 ? emptyQuery : { all: ps };
  },
  None: (_, l, __) => {
    const ps = l.eval();
    return ps.length === 0 ? emptyQuery : { none: ps };
  },
  Sequential: (_, l, __) => {
    const ps = l.eval();
    return ps.length === 0 ? emptyQuery : { sequential: ps };
  },
  Count: (cs, _, exp, __) => {
    let count = cs.eval();
    count.where = exp.eval();
    return count;
  },
  CountSpecifier_countInequality: (_, __, c, n) => {
    let inequality = c.eval();
    switch (inequality) {
      case ">=":
        return { moreThan: parseInt(n.eval()) - 1 };
      case ">":
        return { moreThan: n.eval() };
      case "<=":
        return { lessThan: parseInt(n.eval()) + 1 };
      case "<":
        return { lessThan: n.eval() };
      case ":":
      case "=":
        return { count: n.eval() };
    }
    throw new Error('Unexpected inequality ' + inequality);
  },
  CountSpecifier_countN: (_, n) => { return { count: n.eval() }; },
  CountSpecifier_count3: (_) => {return {count: 3}; },
  CountSpecifier_count2: (_) => {return {count: 2}; },
  CountSpecifier_count1: (_) => {return {count: 1}; },
  linkExp: (l, colon, r) => {
    const ps = r.eval();
    return ps.length === 0 ? emptyQuery : {link: ps };
  },
  Exp: orConjunction,
  NestedExp: evalSelf,
  NestedExp_paren: (_, p, __) => p.eval(),
  NestedExp_not: evalNot,
  OrPart: andConjunction,
  AndPart_fragment: evalSelf,
  Fragment: evalSelf,
  Fragment_not: evalNot,
  browserName: (browser) => {
    return browser.sourceString.toUpperCase();
  },
  statusLiteral: (status) => {
    return status.sourceString.toUpperCase() === 'MISSING'
        ? 'UNKNOWN'
        : status.sourceString.toUpperCase();
  },
  statusExp_eq: (l, colon, r) => {
    return { status: r.eval() };
  },
  statusExp_product_eq: (l, colon, r) => {
    return {
      product: l.sourceString.toLowerCase(),
      status: r.eval(),
    };
  },
  statusExp_neq: (l, colonBang, r) => {
    return { status: {not: r.eval() } };
  },
  statusExp_product_neq: (l, colonBang, r) => {
    return {
      product: l.sourceString.toLowerCase(),
      status: {not: r.eval()},
    };
  },
  isExp: (l, colon, r) => {
    return { is: r.eval() };
  },
  triagedExp: (l, colon, r) => {
    const ps = r.eval();
    if (ps.length === 0) {
      return emptyQuery;
    }
    // Test-level issues are represented on the backend as an empty product.
    return { triaged: ps.toLowerCase().replace('test-issue', '') };
  },
  labelExp: (l, colon, r) => {
    const ps = r.eval();
    return ps.length === 0 ? emptyQuery : {label: ps };
  },
  webFeatureExp: (l, colon, r) => {
    const ps = r.eval();
    return ps.length === 0 ? emptyQuery : {feature: ps };
  },
  subtestExp: (l, colon, r) => {
    return { subtest: r.eval() };
  },
  pathExp: (l, colon, r) => {
    return { path: r.eval() };
  },
  patternExp: (p) => {
    return { pattern: p.eval() };
  },
  nameFragment_basic: (p) => {
    return p.sourceString;
  },
  nameFragment_quoted: (_, chars,  __) => {
    return chars.sourceString;
  },
  backslash: (v) => '\\',
  quotemark: (v) => '"',
  number: (v) => parseInt(v.sourceString),
});
/* eslint-enable */

const QUERY_DEBOUNCE_ID = Symbol('query_debounce_timeout');

class TestSearch extends WPTFlags(PolymerElement) {
  static get template() {
    return html`
    <style>
      input.query {
        font-size: 16px;
        display: block;
        padding: 0.5em 0;
        width: 100%;
      }
      .help {
        float: right;
      }
    </style>

    <div>
      <input class="query" list="query-list" aria-label="Search test files"
             value="{{ queryInput::input }}" placeholder="[[placeholder]]"
             onchange="[[onChange]]" onkeyup="[[onKeyUp]]" onkeydown="[[onKeyDown]]" onfocus="[[onFocus]]" onblur="[[onBlur]]">
      <span class="help">
        For information on the search syntax, <a href="https://github.com/web-platform-tests/wpt.fyi/blob/main/api/query/README.md">view the search documentation</a>
      </span>

      <!-- TODO(markdittmer): Static id will break multiple search components. -->
      <datalist id="query-list"></datalist>
      <paper-tooltip position="top" manual-mode="true">
        Press &lt;Enter&gt; to commit query
      </paper-tooltip>
    </div>
`;
  }

  static get QUERY_GRAMMAR() {
    return QUERY_GRAMMAR;
  }
  static get QUERY_SEMANTICS() {
    return QUERY_SEMANTICS;
  }
  static get is() {
    return 'test-search';
  }
  static get properties() {
    return {
      placeholder: {
        type: String,
        value: 'Search test files, like \'cors/allow-headers.htm\', then press <Enter>',
      },
      // Query input string
      queryInput: {
        type: String,
        notify: true,
        observer: 'queryInputChanged'
      },
      // Debounced + normalized query string.
      query: {
        type: String,
        notify: true,
        observer: 'queryUpdated',
      },
      structuredQuery: {
        type: Object,
        notify: true,
      },
      results: {
        type: Array,
        notify: true,
      },
      testPaths: Set,
      onKeyUp: Function,
      onChange: Function,
      onFocus: Function,
      onBlur: Function,
    };
  }

  constructor() {
    super();

    this.onChange = this.handleChange.bind(this);
    this.onFocus = this.handleFocus.bind(this);
    this.onBlur = this.handleBlur.bind(this);
    this.onKeyUp = this.handleKeyUp.bind(this);
    this.onKeyDown = this.handleKeyDown.bind(this);
  }

  ready() {
    super.ready();
    this._createMethodObserver('updateDatalist(query, testPaths)');
    this.queryInput = this.query;
  }

  queryUpdated(query) {
    this.queryInput = query;
    if (this.structuredQueries) {
      if (!query) {
        this.structuredQuery = null;
      } else {
        try {
          this.structuredQuery = Object.freeze(this.parseAndInterpretQuery(query));
        // eslint-disable-next-line no-unused-vars
        } catch (err) {
          // TODO: Handle query parse/interpret error.
        }
      }
    }
  }

  parseAndInterpretQuery(query) {
    const p = QUERY_GRAMMAR.match(query);
    if (!p.succeeded()) {
      throw new Error(`Failed to parse query: ${query}`);
    }

    return QUERY_SEMANTICS(p).eval();
  }

  updateDatalist(query, paths) {
    const datalist = this.shadowRoot.querySelector('datalist');
    datalist.innerHTML = '';
    for (const atomPrefix of Object.keys(atoms)) {
      if (!query || atomPrefix.startsWith(query)) {
        const option = document.createElement('option');
        option.setAttribute('value', atomPrefix + ':');
        option.setAttribute('atom', atomPrefix);
        datalist.appendChild(option);
      } else if (query) {
        for (const value of atoms[atomPrefix].map(v => `${atomPrefix}:${v}`)) {
          if (value.startsWith(query)) {
            const option = document.createElement('option');
            option.setAttribute('value', value);
            option.setAttribute('atom', value);
            datalist.appendChild(option);
          }
        }
      }
    }
    if (paths) {
      let matches = Array.from(paths);
      if (query) {
        matches = matches
          .filter(p => p.toLowerCase().includes(query))
          .sort((p1, p2) => p1.indexOf(query) - p2.indexOf(query));
      }
      for (const match of matches.slice(0, 10 - datalist.children.length)) {
        const option = document.createElement('option');
        option.setAttribute('value', match);
        datalist.appendChild(option);
      }
    }
  }

  queryInputChanged(_, oldQuery) {
    // Debounce first initialization.
    if (typeof(oldQuery) === 'undefined') {
      return;
    }
    if (this[QUERY_DEBOUNCE_ID]) {
      window.clearTimeout(this[QUERY_DEBOUNCE_ID]);
    }
    this[QUERY_DEBOUNCE_ID] = window.setTimeout(this.latchQuery.bind(this), 500);
  }

  latchQuery() {
    this.query = (this.queryInput || '').toLowerCase();
  }

  commitQuery() {
    this.query = this.queryInput;
    this.dispatchEvent(new CustomEvent('commit', {
      detail: {
        query: this.query,
        structuredQuery: this.structuredQuery,
      },
    }));
    this.shadowRoot.querySelector('.query').blur();
  }

  handleKeyDown(e) {
    // Prevent tab key navigation on search bar.
    if (e.keyCode === 9) {
      e.preventDefault();
      return false;
    }
  }

  handleKeyUp(e) {
    // Commit when enter key was pressed
    if (e.keyCode === 13) {
      this.commitQuery();
    }
  }

  handleChange(e) {
    const opts = Array.from(this.shadowRoot.querySelectorAll('option'));
    if (opts.length === 0) {
      return;
    }

    const path = e.target.value;
    const autocompleteSelection =
      opts.find(o => o.getAttribute('value').toLowerCase().includes(path.toLowerCase()));
    if (autocompleteSelection) {
      if (autocompleteSelection.getAttribute('atom')) {
        return;
      }
      if (autocompleteSelection.value.toLowerCase() === path.toLowerCase()) {
        this.dispatchEvent(new CustomEvent('autocomplete', {
          detail: {path: autocompleteSelection.value},
        }));
        this.shadowRoot.querySelector('.query').blur();
      }
    }
  }

  handleFocus() {
    this.shadowRoot.querySelector('paper-tooltip').show();
  }

  handleBlur() {
    this.shadowRoot.querySelector('paper-tooltip').hide();
  }

  clear() {
    this.query = '';
    this.queryInput = '';
  }
}
window.customElements.define(TestSearch.is, TestSearch);

export { TestSearch };

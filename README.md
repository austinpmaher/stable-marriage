Solution to the stable marriage problem in Go using the Gale-Shapley algorithm.

    function stableMatching {
        Initialize all m â M and w â W to free
        while â free man m who still has a woman w to propose to {
            w = first woman on mâs list to whom m has not yet proposed
            if w is free
                (m, w) become engaged
            else some pair (m', w) already exists
                if w prefers m to m'
                    m' becomes free
                   (m, w) become engaged 
                else
                   (m', w) remain engaged
        }
    }

pseudocode from [here](https://towardsdatascience.com/gale-shapley-algorithm-simply-explained-caa344e643c2)

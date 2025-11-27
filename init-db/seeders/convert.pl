#!/usr/bin/perl
use strict;
use warnings;

my %digits = (
    '0' => '०', '1' => '१', '2' => '२', '3' => '३', '4' => '४',
    '5' => '५', '6' => '६', '7' => '७', '8' => '८', '9' => '९'
);

while (<>) {
    # Only process lines that have a colon followed by opening quote
    if (/":\s*"/) {
        # Find the ": " separator
        if (/^(.*?":\ ")(.*?)(".*?)$/) {
            my ($prefix, $value, $suffix) = ($1, $2, $3);
            # Replace digits in value only
            foreach my $digit (keys %digits) {
                $value =~ s/$digit/$digits{$digit}/g;
            }
            $_ = $prefix . $value . $suffix . "\n";
        }
    }
    print;
}
